package callback

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/config"
	"github.com/sg3t41/api/model"
	"github.com/sg3t41/api/pkg/redis"
	"github.com/sg3t41/api/pkg/util"
	"github.com/sg3t41/api/pkg/util/jwt"
)

func Get(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not found"})
		return
	}

	// アクセストークンを取得
	githubAccessToken, err := getAccessToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error_accesstoken": err.Error()})
		return
	}

	user, err := getUserInfo(githubAccessToken)
	if err != nil {
		fmt.Println(err)
	}

	/************************************
	* Store github user data to Postgres
	*************************************/
	userExists, err := model.GetRecords("users", "github_id = $1", user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userID string
	if len(userExists) > 0 {
		// ユーザーが存在する場合、更新する
		q := `
		UPDATE users 
		SET
			username=$1,
			email=$2, 
			avatar_url=$3, 
			profile_url=$4, 
			full_name=$5 
		WHERE 
			github_id=$6
		`
		updatedUserID, err := model.UpdateRecord(q, user.Login, user.Email, user.AvatarURL, user.URL, user.Name, user.ID)
		if err != nil {
			fmt.Println(q)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userID = updatedUserID
	} else {
		// ユーザーが存在しない場合、新規作成
		q := `
		INSERT INTO users (
			github_id, 
			username, 
			email, 
			avatar_url,
			profile_url, 
			full_name
		) VALUES ($1, $2, $3, $4, $5, $6)`
		insertedUserID, err := model.CreateRecord(q, user.ID, user.Login, user.Email, user.AvatarURL,
			user.URL, user.Name)
		if err != nil {
			fmt.Println(q)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userID = insertedUserID
	}

	/***************
	 * Create JWT
	 ***************/
	sessionID, err := util.Rand(32)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	jwtToken, err := jwt.GenerateToken(sessionID, userID, user.Login, user.AvatarURL)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	/***************************
	* Store token to Postgres
	****************************/
	refreshToken, err := util.Rand(32)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	q := `
	INSERT INTO user_tokens (
		user_id, 
		access_token,
		refresh_token, 
		expires_at, 
		refresh_expires_at,
		is_revoked
	) 
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = model.CreateRecord(q, userID, jwtToken, refreshToken, time.Now().Add(1*time.Hour), time.Now().Add(7*24*time.Hour), true)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	/************************************************
	 * Store session data and access token to Redis
	 ************************************************/
	oneDay := time.Hour * 24

	// TODO: impl async
	{
		/* Store tokens */
		e := oneDay * 3
		k := fmt.Sprintf("user_token:%s", userID)
		v := map[string]interface{}{
			"access_token":  jwtToken,
			"refresh_token": refreshToken,
		}

		if err := redis.HSet(c, k, v); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if err := redis.Expire(c, k, e); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	{
		/* Store session data */
		k := fmt.Sprintf("session:%s", sessionID)
		v := map[string]interface{}{
			"github_access_token": githubAccessToken,
			"github_id":           user.ID,
			"profile_url":         user.URL,
			"github_name":         user.Login,
			"email":               user.Email,
			"avatar_url":          user.AvatarURL,
			"full_name":           user.Name,
		}

		if err := redis.HSet(c, k, v); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.Redirect(http.StatusFound, "http://localhost:3000/api/set-cookie?token="+jwtToken)
}

func getAccessToken(code string) (string, error) {
	clientID := config.OAuthSetting.GithubClientID
	clientSecret := config.OAuthSetting.GithubClientSecret
	url := "https://github.com/login/oauth/access_token"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Add("client_id", clientID)
	q.Add("client_secret", clientSecret)
	q.Add("code", code)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get access token: %s", resp.Status)
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.AccessToken, nil
}

type User struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"html_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

func getUserInfo(token string) (User, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("failed to get user info: %s", resp.Status)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}

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
	linkedInAccessToken, err := getAccessToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error_accesstoken": err.Error()})
		return
	}

	fmt.Println("access-token:::")
	fmt.Println(linkedInAccessToken)

	user, err := getUserInfo(linkedInAccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error_userinfo": err.Error()})
		return
	}

	/************************************
	* Store LinkedIn user data to Postgres
	*************************************/
	userExists, err := model.GetRecords("users", "linkedin_id = $1", user.ID)
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
			linkedin_id=$6
		`
		updatedUserID, err := model.UpdateRecord(q, user.Name, user.Email, user.AvatarURL, "", user.Name, user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userID = updatedUserID
	} else {
		// ユーザーが存在しない場合、新規作成
		q := `INSERT INTO users (
	    		linkedin_id, 
	    		username, 
	    		email, 
	    		avatar_url,
	    		profile_url, 
	    		full_name
	    	) VALUES ($1, $2, $3, $4, $5, $6)`
		insertedUserID, err := model.CreateRecord(q, user.ID, user.Name, user.Email, user.AvatarURL, "", user.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userID = insertedUserID
	}

	// Store LinkedIn Oauth Token to Postgres
	{
		q := `INSERT INTO oauth_tokens (
			user_id,
			provider,
			access_token
		) VALUES ($1, $2, $3)`
		_, err := model.CreateRecord(q, userID, "linkedin", linkedInAccessToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	/***************
	 * Create JWT
	 ***************/
	sessionID, err := util.Rand(32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	jwtToken, err := jwt.GenerateToken(sessionID, userID, user.Name, user.AvatarURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	/***************************
	* Store token to Postgres
	****************************/
	refreshToken, err := util.Rand(32)
	if err != nil {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	/************************************************
	 * Store session data and access token to Redis
	 ************************************************/
	oneDay := time.Hour * 24

	// Store tokens
	e := oneDay * 3
	k := fmt.Sprintf("user_token:%s", userID)
	v := map[string]interface{}{
		"access_token":  jwtToken,
		"refresh_token": refreshToken,
	}

	if err := redis.HSet(c, k, v); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := redis.Expire(c, k, e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Store session data
	k = fmt.Sprintf("session:%s", sessionID)
	v = map[string]interface{}{
		"linkedin_access_token": linkedInAccessToken,
		"linkedin_id":           user.ID,
		//	"profile_url":           user.ProfileURL,
		//	"linkedin_name":         user.LocalName,
		//	"email":                 user.Email,
		//	"avatar_url":            user.Picture,
		//	"full_name":             user.FullName,
	}

	if err := redis.HSet(c, k, v); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Redirect(http.StatusFound, "http://localhost:3000/api/set-cookie?token="+jwtToken)
}

func getAccessToken(code string) (string, error) {
	clientID := config.OAuthSetting.LinkedInClientID
	clientSecret := config.OAuthSetting.LinkedInClientSecret
	url := "https://www.linkedin.com/oauth/v2/accessToken"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Add("grant_type", "authorization_code")
	q.Add("client_id", clientID)
	q.Add("client_secret", clientSecret)
	q.Add("code", code)
	q.Add("redirect_uri", "http://localhost:8080/api/auth/linkedin/callback")
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
	ID        string `json:"sub"`
	Name      string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	AvatarURL string `json:"picture"`
	// Locale        string `json:"locale"`
	Email         string `json:"email"`
	EmailVelified bool   `json:"email_verified"`
}

func getUserInfo(token string) (User, error) {
	req, err := http.NewRequest("GET", "https://api.linkedin.com/v2/userinfo", nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("failed to get user info: %s", resp.Status)
	}

	fmt.Println("BODYYYYYYYYYYY")
	fmt.Println(resp.Body)

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return User{}, err
	}

	fmt.Println("********************")
	fmt.Println(user)
	fmt.Println("*********************")

	return user, nil
}

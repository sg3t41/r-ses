package callback

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/config"
)

func Get(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not found"})
		return
	}

	// アクセストークンを取得
	token, err := getAccessToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := getUserInfo(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	c.Redirect(http.StatusFound, "http://localhost:3000?token="+token)
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
	// 必要に応じて他のフィールドを追加
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

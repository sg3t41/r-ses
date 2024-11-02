package session

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/model"
	"github.com/sg3t41/api/pkg/redis"
)

func LoadGithubAccessToken(c *gin.Context) (string, error) {
	// Redis
	{
		sessionID, _ := c.Get("sessionID")
		key := fmt.Sprintf("session:%s", sessionID)
		token, isExist, err := redis.HGet(c, key, "github_access_token")
		if err != nil {
			fmt.Println(err)
		}

		if isExist && token != "" {
			return token, nil
		}
	}

	{
		userID, _ := c.Get("userID")
		type Columns struct {
			Token string `db:"access_token"`
		}
		result, err := model.GetRecords2[Columns]("oauth_tokens", []string{"access_token"}, "user_id = $1 AND provider = $2", userID, "github")
		if err != nil {
			return "", err
		}
		token := result[0].Token
		return token, nil
	}
}

package login

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/config"
)

func Get(c *gin.Context) {
	clientID := config.OAuthSetting.GithubClientID
	redirectURI := config.OAuthSetting.GithubRedirectURL
	authPageFmt := "https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s"
	fmt.Println(clientID)
	fmt.Println(config.OAuthSetting.GithubClientSecret)
	url := fmt.Sprintf(authPageFmt, clientID, redirectURI)
	c.Redirect(http.StatusFound, url)
}

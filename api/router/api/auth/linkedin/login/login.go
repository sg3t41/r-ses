package login

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/config"
)

func Get(c *gin.Context) {
	clientID := config.OAuthSetting.LinkedInClientID
	redirectURI := config.OAuthSetting.LinkedInRedirectURL
	authPageFmt := "https://www.linkedin.com/oauth/v2/authorization?response_type=code&client_id=%s&redirect_uri=%s&scope=openid%%20profile%%20email"
	url := fmt.Sprintf(authPageFmt, clientID, redirectURI)
	c.Redirect(http.StatusFound, url)
}

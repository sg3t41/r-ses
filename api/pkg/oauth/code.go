package oauth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) (string, error) {
	code := c.Query("code")
	if code == "" {
		err := fmt.Errorf("failed to get code")
		return "", err
	}
	return code, nil
}

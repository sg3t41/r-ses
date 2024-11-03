package portfolio

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	PortfolioRepoIDs []string `json:"portfolioRepoIds"`
}

func Post(c *gin.Context) {
	var req PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// エラーハンドリング
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	fmt.Println("Received IDs:", req.PortfolioRepoIDs)

	// db保存と項目を考える TODO

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

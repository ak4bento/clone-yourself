package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type AskRequest struct {
	Question string `json:"question"`
}

type AskResponse struct {
	Answer string `json:"answer"`
}

func StartAPI() {
	r := gin.Default()

	r.POST("/ask", func(c *gin.Context) {
		var req AskRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// TODO: Integrasi dengan memory dan analyzer
		c.JSON(http.StatusOK, AskResponse{Answer: "Docker itu ibarat kontainer di pelabuhan..."})
	})

	r.Run() // :8080
}

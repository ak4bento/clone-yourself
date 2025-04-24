package server

import (
	"github.com/ak4bento/clone-yourself/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
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

		answer, _ := core.FindRelevantKnowledge(req.Question)
		core.LearnFromInteraction(req.Question, answer)

		c.JSON(http.StatusOK, AskResponse{Answer: answer})
	})

	r.Run() // :8080
}

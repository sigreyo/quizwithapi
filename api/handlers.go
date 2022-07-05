package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

func getHighScore(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, answers)
}

func postAnswers(c *gin.Context) {
	var newScore Answer

	if err := c.BindJSON(&newScore); err != nil {
		return
	}

	answers = append(answers, newScore)
	c.IndentedJSON(http.StatusCreated, newScore)
}

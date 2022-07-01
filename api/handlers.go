package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

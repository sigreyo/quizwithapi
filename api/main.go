package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/questions/", getAllQuestions)

	router.GET("/api/answers/", getHighScore)
	router.POST("/api/answers/", postAnswers)

	router.Run("localhost:8080")
}

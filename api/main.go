package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/questions", getAllQuestions)

	router.Run("localhost:8080")
}

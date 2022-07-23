package main

import (
	"os"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.GET("/entries", nil)

	router.POST("/entry/create", nil)
	router.GET("/entry/:id", nil)
	router.PUT("entry/update/:id", nil)
	router.DELETE("/entry/delete/:id", nil)

	router.GET("/ingredient/:ingredient", nil)
	router.PUT("/ingredient/:ingredient", nil)

	router.Run(":", port)
}

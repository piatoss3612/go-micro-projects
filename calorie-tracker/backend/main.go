package main

import (
	"backend/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
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

	router.GET("/entries", routes.GetEntries)

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entry/:id", routes.GetEntryById)
	router.PUT("entry/update/:id", routes.UpdateEntry)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)

	router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredient)

	if err := router.Run(":" + port); err != nil {
		log.Panic(err)
	}
}

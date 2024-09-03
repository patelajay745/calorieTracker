package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/patelajay745/calories-tracker/controllers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", controllers.AddEntry)
	router.GET("/entries", controllers.GetEntries)
	router.GET("/entry/:id/", controllers.GetEntryByID)
	router.GET("/ingredient/:ingredient", controllers.GetEntriesByIngredient)

	router.PUT("/entry/update/:id", controllers.UpdateEntry)
	router.PUT("/ingredient/update/:id", controllers.UpdateIngredient)

	router.DELETE("/entry/delete/:id", controllers.DeleteEntry)

	router.Run(":" + port)

}

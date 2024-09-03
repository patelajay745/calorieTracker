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

	router.POST("/entry", controllers.AddEntry)         //tested:okay
	router.GET("/entry", controllers.GetEntries)        //tested:okay
	router.GET("/entry/:id/", controllers.GetEntryByID) //tested:okay
	router.GET("/ingredient/:ingredient", controllers.GetEntriesByIngredient)

	router.PUT("/entry/:id", controllers.UpdateEntry) 
	router.PUT("/ingredient/:id", controllers.UpdateIngredient)

	router.DELETE("/entry/:id", controllers.DeleteEntry) //tested:okay

	router.Run(":" + port)

}

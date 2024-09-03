package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddEntry(c *gin.Context) {

}

func GetEntries(c *gin.Context) {

}

func EntryByID(c *gin.Context) {

}

func GetEntriesByIngredient(c *gin.Context) {

}

func UpdateEntry(c *gin.Context) {

}
func UpdateIngredient(c *gin.Context) {

}
func DeleteEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	primitive.ObjectIDFromHex(entryID)

	 

}

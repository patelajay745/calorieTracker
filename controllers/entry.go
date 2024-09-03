package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/patelajay745/calories-tracker/config"
	"github.com/patelajay745/calories-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var entryCollection *mongo.Collection = config.GetCollection(config.Client, "calories")

func AddEntry(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	var entry models.Entry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	if validationErr := validate.Struct(entry); validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		return
	}
	entry.ID = primitive.NewObjectID()

	_, err := entryCollection.InsertOne(ctx, entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, entry)
}

func GetEntries(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	if entries == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No entries in database"})
		return
	}

	c.JSON(http.StatusOK, entries)
}

func GetEntryByID(c *gin.Context) {
	EntryId := c.Params.ByName("id")
	fmt.Println(EntryId)

	docID, _ := primitive.ObjectIDFromHex(EntryId)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	var entry bson.M
	cursor := entryCollection.FindOne(ctx, bson.M{"_id": docID})

	defer cancel()
	if err := cursor.Decode(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func GetEntriesByIngredient(c *gin.Context) {
	ingredientId := c.Params.ByName("id")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{"ingredients": ingredientId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entries)
}

func UpdateEntry(c *gin.Context) {
	EntryId := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(EntryId)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var newEntry models.Entry
	if err := c.BindJSON(&newEntry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if validationErr := validate.Struct(newEntry); validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		return
	}

	var oldEntry models.Entry
	entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&oldEntry)

	if newEntry.Dish != nil {
		oldEntry.Dish = newEntry.Dish
	}
	if newEntry.Fat != nil {
		oldEntry.Fat = newEntry.Fat
	}
	if newEntry.Ingredients != nil {
		oldEntry.Ingredients = newEntry.Ingredients
	}
	if newEntry.Calories != nil {
		oldEntry.Calories = newEntry.Calories
	}

	updated, err := entryCollection.ReplaceOne(ctx, bson.M{"_id": docID}, oldEntry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated successfully", "count": updated.ModifiedCount})

}
func UpdateIngredient(c *gin.Context) {

	EntryId := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(EntryId)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	type Ingredient struct {
		Ingredients *string `json:"ingredients"`
	}

	var ingredient Ingredient

	if err := c.BindJSON(&ingredient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updated, err := entryCollection.UpdateOne(ctx, bson.M{"_id": docID},
		bson.D{{"$set", bson.D{{"ingredients", ingredient.Ingredients}}}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated successfully", "count": updated.ModifiedCount})

}
func DeleteEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docId, _ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully", "count": result.DeletedCount})

}

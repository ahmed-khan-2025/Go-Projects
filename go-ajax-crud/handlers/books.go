package handlers

import (
	"context"
	"net/http"

	"go-ajax-crud/db"
	"go-ajax-crud/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(c *gin.Context) {
	cursor, _ := db.BookCollection.Find(context.Background(), bson.M{})

	var books []models.Book
	cursor.All(context.Background(), &books)

	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var b models.Book
	c.ShouldBindJSON(&b)

	db.BookCollection.InsertOne(context.Background(), b)

	c.JSON(200, gin.H{"message": "created"})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var b models.Book
	c.ShouldBindJSON(&b)

	objID, _ := primitive.ObjectIDFromHex(id)

	db.BookCollection.UpdateOne(context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{
			"title":  b.Title,
			"author": b.Author,
		}})

	c.JSON(200, gin.H{"message": "updated"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	db.BookCollection.DeleteOne(context.Background(), bson.M{"_id": objID})

	c.JSON(200, gin.H{"message": "deleted"})
}
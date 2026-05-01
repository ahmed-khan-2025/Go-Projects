package handlers

import (
	"context"
	"net/http"

	"go-ajax-crud/db"
	"go-ajax-crud/models"
	"go-ajax-crud/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var u models.User
	c.ShouldBindJSON(&u)

	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hash)

	db.UserCollection.InsertOne(context.Background(), u)

	c.JSON(http.StatusOK, gin.H{"message": "registered"})
}

func Login(c *gin.Context) {
	var input models.User
	var dbUser models.User

	c.ShouldBindJSON(&input)

	err := db.UserCollection.FindOne(context.Background(),
		bson.M{"username": input.Username}).Decode(&dbUser)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
		return
	}

	token, _ := utils.GenerateToken(dbUser.Username)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
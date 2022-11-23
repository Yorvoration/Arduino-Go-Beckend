package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb+srv://root:1234@cluster0.ik76ncs.mongodb.net/?retryWrites=true&w=majority"

type User struct {
	OnLamp bool `json:"onLamp"`
	OnMator bool `json:"onMator"`
}

func main() {
	router := gin.Default()
	router.POST("/on", on)
	router.Run()

	fmt.Println("Server is running on port 8080")
}

func on(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}
	collection := client.Database("ESP8266MOD").Collection("on")
	var user User
	err = c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
	}
	insertResult, err := collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	c.JSON(200, gin.H{
		"message": "success",
	})
	
}
	



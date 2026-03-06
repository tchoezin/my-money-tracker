package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("Hello, World!")
	//creates a defualt Gin router
	r := gin.Default()

	//Define a simple GET route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Define a GET route with a parameter
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
			"name":    "Choezin",
		})
	})

	var user struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required"`
		}

	//Define a POST route
	r.POST("/user", func(c *gin.Context) {
		
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.JSON(http.StatusCreated, gin.H{
				"message": "User created",
				"user":    user,
			})
		}
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"email": 
		})
	})

	//Start server on port 8080
	r.Run(":8080")

}

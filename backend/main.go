package main

import (
	"my-money-tracker/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	transactions := r.Group("/transactions")
	{
		transactions.GET("", handlers.GetTransactions)
		transactions.GET("/:id", handlers.GetTransaction)
		transactions.POST("", handlers.CreateTransaction)
		transactions.PUT("/:id", handlers.UpdateTransaction)
		transactions.DELETE("/:id", handlers.DeleteTransaction)
	}

	//Start server on port 8080
	r.Run(":8080")

}

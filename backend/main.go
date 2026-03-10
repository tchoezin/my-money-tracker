package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//transactions := r.Group("/transactions")

	//Start server on port 8080
	r.Run(":8080")

}

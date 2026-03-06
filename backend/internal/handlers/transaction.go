package handlers

import (
	"my-money-tracker/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTransactions handles GET /transactions
func GetTransactions(c *gin.Context) {
	transactions := models.GetAllTransactions()
	c.JSON(http.StatusOK, transactions)
}

// GetTransaction handles GET /transactions/:id
func GetTransaction(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction, err := models.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, transaction)
}

// CreateTransaction handles POST /transactions
func CreateTransaction(c *gin.Context) {

}

// UpdateTransaction handles PUT /transactions/:id
func UpdateTransaction(c *gin.Context) {

}

// DeleteTransaction handles DELETE /transactions/:id
func DeleteTransaction(c *gin.Context) {
	strID := c.Param("id")
}

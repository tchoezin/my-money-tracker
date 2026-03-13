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
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// CreateTransaction handles POST /transactions
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validation
	if transaction.Type != "income" && transaction.Type != "expense" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type must be 'income' or 'expense'"})
		return
	}

	if transaction.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be positive"})
		return
	}

	created := models.AddTransaction(transaction)

	c.JSON(http.StatusCreated, created)
}

// UpdateTransaction handles PUT /transactions/:id
func UpdateTransaction(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedTransaction models.Transaction
	if err = c.ShouldBindJSON(&updatedTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTransaction, err := models.UpdateTransaction(id, updatedTransaction)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTransaction)
}

// DeleteTransaction handles DELETE /transactions/:id
func DeleteTransaction(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteTransaction(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

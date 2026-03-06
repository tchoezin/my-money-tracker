package models

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID          int
	Amount      float64
	Date        time.Time
	Type        string // "income" or "expense"
	Description string
	Category    string
}

// in-memory data store
var transactions = []Transaction{}
var nextID = 1

// Get all transactions
func GetAllTransactions() []Transaction {
	return transactions
}

// Add a transaction
func AddTransaction(t Transaction) {
	t.ID = nextID
	nextID++
	t.Date = time.Now()

	transactions = append(transactions, t)
}

// Remove a transaction by id
func DeleteTransaction(id int) error {
	for i, transaction := range transactions {
		if transaction.ID == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("failed finding transaction id: %v for deletion", id)
}

// Get a transaction by id
func GetTransaction(id int) (Transaction, error) {
	for _, transaction := range transactions {
		if transaction.ID == id {
			return transaction, nil
		}
	}
	return Transaction{}, fmt.Errorf("failed getting transaction id: %v", id)
}

// Update a transaction by id
func UpdateTransaction(id int, updatedTransaction Transaction) error {
	for i, transaction := range transactions {
		if transaction.ID == id {
			transactions[i] = updatedTransaction
			return nil
		}
	}
	return fmt.Errorf("failed to find transaction id: %v for update", id)
}

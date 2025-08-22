package main

import "time"

type Transaction struct {
	Id       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string // income or expense
}

// BudgetTracker Struct to manage transactions
type BudgetTracker struct {
	transactions []Transaction
	nextId       int
}

// FinancialRecord Interface for common behavior
type FinancialRecord interface {
	GetAmount() float64
	GetType() string
}

// GetAmount Implement interface methods for Transaction struct
func (t Transaction) GetAmount() float64 {
	return t.Amount
}

// GetType Implement interface methods for Transaction struct
func (t Transaction) GetType() string {
	return t.Type
}

// Add a new Transaction to the budget tracker
func (bt *BudgetTracker) AddTransaction(amount float64, category string, tType string) {
	newTransaction := Transaction{
		Id:       bt.nextId,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     tType,
	}
}

func main() {

}

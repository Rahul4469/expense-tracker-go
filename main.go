package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Transaction entity
type Transaction struct {
	Id       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string // income" or "expense"
}

// BudgetTracker entity
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

// AddTransaction Function to add a new Transaction to the budget tracker
// by creating a new instance of Transaction
func (bt *BudgetTracker) AddTransaction(amount float64, category string, tType string) {
	newTransaction := Transaction{
		Id:       bt.nextId,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     tType,
	}
	//transaction belongs to the BudgetTracker entity
	bt.transactions = append(bt.transactions, newTransaction)
	bt.nextId++
}

func (bt BudgetTracker) DisplayTransactions() {
	fmt.Println("ID\tAmount\tCategory\tDate\t\t\tType")
	for _, transaction := range bt.transactions {
		fmt.Printf("%d\t%.2f\t%s\t\t%s\t%s\n",
			transaction.Id, transaction.Amount, transaction.Category,
			transaction.Date.Format("2006-01-02"), transaction.Type)
	}
}

// CalculateTotal Get total income / expense
func (bt BudgetTracker) CalculateTotal(tType string) float64 {
	var total float64
	for _, transaction := range bt.transactions {
		if transaction.Type == tType {
			total += transaction.Amount
		}
	}
	return total
}

// SaveToCsv saves the transactions to a CSV file.
func (bt BudgetTracker) SaveToCsv(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Amount", "Category", "Date", "Type"})

	//write data
	for _, t := range bt.transactions {
		record := []string{
			strconv.Itoa(t.Id),
			fmt.Sprintf("%.2f", t.Amount),
			t.Category,
			t.Date.Format("2006-01-02"),
			t.Type,
		}
		writer.Write(record)
	}
	fmt.Println("Transaction saved to", filename)
	return nil
}

func main() {
	//Instatiation of BudgetTracker entity
	bt := BudgetTracker{}
	for {
		fmt.Println("\n---Personal Budget Tracker---")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Display Transaction")
		fmt.Println("3. Show Total Income")
		fmt.Println("4. Show Total Expenses")
		fmt.Println("5. Save Transaction to CSV file")
		fmt.Println("6. Exit")
		fmt.Println("Choose an option:")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter Amount:")
			var amount float64
			fmt.Scanln(&amount)

			fmt.Println("Enter Category:")
			var category string
			fmt.Scanln(&category)

			fmt.Println("Enter Type (Income/Expense):")
			var tType string
			fmt.Scanln(&tType)

			bt.AddTransaction(amount, category, tType)
			fmt.Println("Transaction Added!")

		case 2:
			bt.DisplayTransactions()
		case 3:
			fmt.Println("Total Income: %.2f\n", bt.CalculateTotal("income"))
		case 4:
			fmt.Println("Total Expense: %.2f\n", bt.CalculateTotal("expense"))
		case 5:
			fmt.Println("Enter filename: (eg. transactions.csv) ")
			var filename string
			fmt.Scanln(&filename)
			if err := bt.SaveToCsv(filename); err != nil {
				fmt.Println("Error saving transactions:", err)
			}
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice! Try Again.")

		}
	}
}

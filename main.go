package main

/*
Importing using libraries
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Transaction represents a financial transaction with its details.
// It includes the ID, date, description, category, amount, and type of the transaction.
type Transaction struct {
	Id          int    `json:"id"`          // Unique identifier for the transaction
	Date        string `json:"date"`        // Date of the transaction in string format
	Description string `json:"description"` // Description of the transaction
	Category    string `json:"category"`    // Category of the transaction (e.g., food, salary)
	Amount      int    `json:"amount"`      // Amount of money involved in the transaction
	Type        string `json:"type"`        // Type of transaction (e.g., income or expense)
}

// Data represents the overall financial data for a user in a specific month.
// It includes the user's name, the month of the data, and a list of transactions.
type Data struct {
	User         string        `json:"user"`         // Name of the user
	Month        string        `json:"month"`        // Month for which the data is recorded
	Transactions []Transaction `json:"transactions"` // List of transactions for the month
}

// main is the entry point of the application.
// It reads the financial data from a JSON file, processes the transactions,
// and prints the income, expenses, and remaining balance or excess expenses.
func main() {
	file, err := os.ReadFile("data.json")
	ifERR(err)
	var data Data
	err = json.Unmarshal(file, &data)
	ifERR(err)
	income, outcome := incomeOutcome(data)
	fmt.Printf("Прибыль за %v: %v рублей\n", data.Month, income)
	fmt.Printf("Траты за %v: %v рублей\n", data.Month, outcome)
	if income >= outcome {
		fmt.Printf("Остаток на %v: %v рублей\n", data.Month, income-outcome)
	} else {
		fmt.Printf("Превышение трат за %v: %v рублей\n", data.Month, outcome-income)
	}
	fmt.Println("Самет - черт")
}

// incomeOutcome calculates the total income and expenses from the given data.
// It iterates through the transactions and sums up the amounts based on their type.
// Returns the total income and total expenses.
func incomeOutcome(d Data) (int, int) {
	income, outcome := 0, 0
	for _, v := range d.Transactions {
		if v.Type == "income" {
			income += v.Amount
		}
		if v.Type == "expense" {
			outcome += v.Amount
		}
	}
	return income, outcome
}

// ifERR checks if an error occurred and logs it if so.
// It terminates the program with a fatal log message if an error is present.
func ifERR(err error) {
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

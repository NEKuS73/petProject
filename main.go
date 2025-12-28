package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Transactions struct {
	Id          int    `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
	Type        string `json:"type"`
}
type Data struct {
	User        string         `json:"user"`
	Month       string         `json:"month"`
	Transaction []Transactions `json:"transactions"`
}

func main() {
	file, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	income, outcome := incomeOutcome(data)
	fmt.Printf("Дорогой, %v, за %v года, Вы заработали %v рублей, а потратили %v рублей\n", data.User, data.Month, income, outcome)
	fmt.Printf("Итого в копилку Вы можете отложить %v рублей\n", income-outcome)
}
func incomeOutcome(d Data) (int, int) {
	income, outcome := 0, 0
	for _, v := range d.Transaction {
		if v.Type == "income" {
			income += v.Amount
		}
		if v.Type == "expense" {
			outcome += v.Amount
		}
	}
	return income, outcome
}

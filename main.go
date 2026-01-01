package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Transaction struct {
	Id          int    `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
	Type        string `json:"type"`
}

type Data struct {
	User         string        `json:"user"`
	Month        string        `json:"month"`
	Transactions []Transaction `json:"transactions"`
}

func main() {
	file, err := os.ReadFile("data.json")
	ifERR(err)
	var data Data
	err = json.Unmarshal(file, &data)
	ifERR(err)
	income, outcome := incomeOutcome(data)
	printAnalitics(data, income, outcome)
}

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

func ifERR(err error) {
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}
func printAnalitics(d Data, inc int, out int) {
	fmt.Printf("Прибыль за %v: %v рублей\n", d.Month, inc)
	fmt.Printf("Траты за %v: %v рублей\n", d.Month, out)
	if inc >= out {
		fmt.Printf("Остаток на %v: %v рублей\n", d.Month, inc-out)
	} else {
		fmt.Printf("Превышение трат за %v: %v рублей\n", d.Month, out-inc)
	}
}

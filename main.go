package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var categoryNames = map[string]string{
	// Доходы
	"salary":       "Зарплата",
	"freelance":    "Фриланс",
	"investments":  "Инвестиции",
	"state":        "Гос. выплаты",
	"other_income": "Прочие доходы",
	// Расходы
	"housing":    "Жильё",
	"food":       "Еда",
	"transport":  "Транспорт",
	"health":     "Здоровье",
	"education":  "Образование",
	"obligatory": "Обязательные платежи",
	"other":      "Прочие расходы",
}

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
	checkERR(err)
	var data Data
	err = json.Unmarshal(file, &data)
	checkERR(err)
	incomeTotal, incomeByCat := calculateStats(data, "income")
	expenseTotal, expenseByCat := calculateStats(data, "expense")
	printReport(data.Month, incomeTotal, incomeByCat, expenseTotal, expenseByCat)
}

func checkERR(err error) {
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}
func calculateStats(d Data, transType string) (int, map[string]int) {
	total := 0
	byCategory := make(map[string]int)
	for _, t := range d.Transactions {
		if t.Type == transType {
			total += t.Amount
			byCategory[t.Category] += t.Amount
		}
	}
	return total, byCategory
}
func printReport(month string, incomeTotal int, incomeByCat map[string]int, expenseTotal int, expenseByCat map[string]int) {
	fmt.Printf("📊 Отчёт за %s\n\n", month)
	fmt.Printf("Всего заработано:   +%d₽\n", incomeTotal)
	fmt.Printf("Всего потрачено:    -%d₽\n", expenseTotal)
	fmt.Printf("Итоговый баланс:    %+d₽\n\n", incomeTotal-expenseTotal)

	fmt.Println("ДОХОДЫ по категориям:")
	for cat, amount := range incomeByCat {
		name := categoryNames[cat]
		if name == "" {
			name = cat
		}
		fmt.Printf("  %-20s +%d₽\n", name+":", amount)
	}

	fmt.Println("\nРАСХОДЫ по категориям:")
	for cat, amount := range expenseByCat {
		name := categoryNames[cat]
		if name == "" {
			name = cat
		}
		fmt.Printf("  %-20s -%d₽\n", name+":", amount)
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

	addTransaction(&data)
	saveData("data.json", data)
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
func addTransaction(d *Data) {
	var newTransaction Transaction
	var err error
	newTransaction.Id = getNextID(d.Transactions)
	newTransaction.Date = time.Now().Format(time.DateOnly)
	newTransaction.Type = inputType()
	newTransaction.Category = inputCat()
	fmt.Println("Расскажите подробнее, какое событие хотите добавить")
	newTransaction.Description = input()
	newTransaction.Amount = inputAmount()
	checkERR(err)
	d.Transactions = append(d.Transactions, newTransaction)
	fmt.Printf("Поздравляю, %v, Вы успешно добавили новую транзакцию!\n", d.User)
	fmt.Println("Более детальная информация:")
	fmt.Printf("Id: %v\nДата совершения: %v\nТип транзакции: %v\nСобытие: %v\nКатегория: %v\nСумма: %v₽\n", newTransaction.Id, newTransaction.Date, newTransaction.Type, newTransaction.Description, newTransaction.Category, newTransaction.Amount)
}
func getNextID(transactions []Transaction) int {
	maxID := 0
	for _, t := range transactions {
		if t.Id > maxID {
			maxID = t.Id
		}
	}
	return maxID + 1
}
func input() string {
	reader := bufio.NewReader(os.Stdin)
	inputValue, err := reader.ReadString('\n')
	checkERR(err)
	return strings.TrimSpace(inputValue)
}
func inputAmount() int {
	for {
		fmt.Println("Какова сумма данной операции")
		amountStr := input()
		amount, err := strconv.Atoi(amountStr)
		if err == nil && amount > 0 {
			return amount
		}
		fmt.Println("Ошибка! Введите положительное число.")
	}
}
func inputCat() string {
	for {
		fmt.Println("Укажите, к какой категории относится данная операция")
		category := input()
		for k, v := range categoryNames {
			if category == v || category == k {
				return k
			}
		}
		fmt.Println("Ошибка: укажите существующую категорию")
	}
}
func inputType() string {
	for {
		fmt.Println("Какой тип транзакции хотите добавить? (income/expense)")
		transactionType := input()
		if transactionType == "income" || transactionType == "expense" {
			return transactionType
		}
		fmt.Println("Ошибка: укажите существующий тип транзакции")
	}
}
func saveData(fileName string, d Data) error {
	jsonData, err := json.MarshalIndent(d, "", " ")
	checkERR(err)
	return os.WriteFile(fileName, jsonData, 0644)
}

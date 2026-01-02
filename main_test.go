package main

import (
	"testing"
)

func TestCalculateStats(t *testing.T) {
	testData := Data{
		Transactions: []Transaction{
			{Id: 1, Amount: 1000, Type: "income", Category: "salary"},
			{Id: 2, Amount: 500, Type: "expense", Category: "food"},
			{Id: 3, Amount: 2000, Type: "income", Category: "freelance"},
			{Id: 4, Amount: 300, Type: "expense", Category: "transport"},
		},
	}
	incomeTotal, incomeByCat := calculateStats(testData, "income")
	expenseTotal, expenseByCat := calculateStats(testData, "expense")

	if incomeTotal != 3000 {
		t.Errorf("Неверный подход: ожидалось 3000, получилось %d", incomeTotal)
	}
	if expenseTotal != 800 {
		t.Errorf("Неверный подход: ожидалось 800, получилось %d", expenseTotal)
	}
	if incomeByCat["salary"] != 1000 {
		t.Errorf("Неправильная сумма для salary. Ожидалось: 1000, Получено: %d", incomeByCat["salary"])
	}
	if incomeByCat["freelance"] != 2000 {
		t.Errorf("Неправильная сумма для freelance. Ожидалось: 2000, Получено: %d", incomeByCat["freelance"])
	}
	if expenseByCat["food"] != 500 {
		t.Errorf("Неправильная сумма для salary. Ожидалось: 500, Получено: %d", expenseByCat["food"])
	}
	if expenseByCat["transport"] != 300 {
		t.Errorf("Неправильная сумма для freelance. Ожидалось: 2000, Получено: %d", expenseByCat["transport"])
	}
}

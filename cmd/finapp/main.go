package main

import (
	"fmt"
	"tetsfile/internal/logic"
	"tetsfile/internal/storage"
	"tetsfile/internal/ui"
)

func main() {
	data := storage.LoadData("../../data.json")

	for {
		fmt.Println("\n=== Личный финансовый менеджер ===")
		fmt.Println("1. Показать полный отчёт")
		fmt.Println("2. Добавить транзакцию")
		fmt.Println("3. Просмотреть список транзакций")
		fmt.Println("4. Удалить транзакцию")
		fmt.Println("5. Выйти из программы")

		fmt.Print("Выберите действие: \n")
		choice := ui.Input()
		switch choice {
		case "1":
			incomeTotal, incomeByCat := logic.CalculateStats(data, "income")
			expenseTotal, expenseByCat := logic.CalculateStats(data, "expense")
			ui.PrintReport(data.Month, incomeTotal, incomeByCat, expenseTotal, expenseByCat)
		case "2":
			ui.AddTransaction(&data)
			if err := storage.SaveData("../../data.json", data); err != nil {
				fmt.Printf("КРИТИЧЕСКАЯ ОШИБКА: не удалось сохранить данные: %v\n", err)
			}
		case "3":
			ui.ListTransactions(data.Transactions)
		case "4":
			ui.DeleteTransaction(&data)
			if err := storage.SaveData("../../data.json", data); err != nil {
				fmt.Printf("КРИТИЧЕСКАЯ ОШИБКА: не удалось сохранить данные: %v\n", err)
			}
		case "5":
			fmt.Println("Сохранение данных и выход...")
			if err := storage.SaveData("../../data.json", data); err != nil {
				fmt.Printf("КРИТИЧЕСКАЯ ОШИБКА: не удалось сохранить данные: %v\n", err)
			}
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова!")
		}
	}
}

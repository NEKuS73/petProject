package main

import (
	"fmt"
	"tetsfile/internal/logic"
	"tetsfile/internal/storage"
	"tetsfile/internal/ui"
)

func main() {
	data := storage.LoadData("data.json")

	for {
		fmt.Println("\n=== Личный финансовый менеджер ===")
		fmt.Println("1. Показать полный отчёт")
		fmt.Println("2. Просмотреть список транзакций")
		fmt.Println("3. Добавить транзакцию")
		fmt.Println("4. Редактировать транзакцию")
		fmt.Println("5. Удалить транзакцию")
		fmt.Println("6. Выйти из программы")

		fmt.Print("Выберите действие: \n")
		choice := ui.Input()
		switch choice {
		case "1":
			incomeTotal, incomeByCat := logic.CalculateStats(data, "income")
			expenseTotal, expenseByCat := logic.CalculateStats(data, "expense")
			ui.PrintReport(data.Month, incomeTotal, incomeByCat, expenseTotal, expenseByCat)
		case "3":
			ui.AddTransaction(&data)
			if err := storage.SaveData("data.json", data); err != nil {
				fmt.Printf("КРИТИЧЕСКАЯ ОШИБКА: не удалось сохранить данные: %v\n", err)
			}
		case "2":
			ui.ListTransactions(data.Transactions)
		case "4":
			ui.EditTransaction(&data)
			if err := storage.SaveData("data.json", data); err != nil {
				fmt.Printf("КРИТИЧЕСКАЯ ОШИБКА: не удалось сохранить данные: %v\n", err)
			}
		case "5":
			ui.DeleteTransaction(&data)
			if err := storage.SaveData("data.json", data); err != nil {
				fmt.Printf("КРИТИЧЕСКАЯ ОШИБКА: не удалось сохранить данные: %v\n", err)
			}
		case "6":
			fmt.Println("Сохранение данных и выход...")
			if err := storage.SaveData("data.json", data); err != nil {
				fmt.Printf("КРИТИЧЕСКАЯ ОШИБКА: не удалось сохранить данные: %v\n", err)
			}
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова!")
		}
	}
}

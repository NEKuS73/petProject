package main

import (
	"fmt"
	"tetsfile/internal/logic"
	"tetsfile/internal/models"
	"tetsfile/internal/storage"
	"tetsfile/internal/ui"
)

func main() {
	data := storage.LoadData("data.json")
	access, userId := ui.Authefication(&data)
	if userId == -1 {
		return
	}
	if err := storage.SaveData("data.json", data); err != nil {
		fmt.Printf("КРИТИЧЕСКАЯ ОШИБКА: не удалось сохранить данные: %v\n", err)
		return
	}
	if !access {
		return
	}

	for {
		fmt.Println("\n=== Личный финансовый менеджер ===")
		fmt.Println("1. Транзакции")
		fmt.Println("2. Профиль")
		fmt.Println("0. Выйти из программы")
		fmt.Print("Выберите действие: ")

		choice := ui.Input()

		switch choice {
		case "1":
			manageTransactions(&data, userId)
		case "2":
			manageProfile(&data, &access, &userId)
		case "0":
			handleExit(&data)
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова!")
		}
	}
}

func manageTransactions(data *models.Database, userId int) {
	for {
		fmt.Println("\n=== Управление транзакциями ===")
		fmt.Println("1. Показать полный отчёт")
		fmt.Println("2. Просмотреть список транзакций")
		fmt.Println("3. Добавить транзакцию")
		fmt.Println("4. Редактировать транзакцию")
		fmt.Println("5. Удалить транзакцию")
		fmt.Println("0. Назад в главное меню")
		fmt.Print("Выберите действие: ")

		transactionChoice := ui.Input()

		switch transactionChoice {
		case "1":
			incomeTotal, incomeByCat := logic.CalculateStats(*data, "income", userId)
			expenseTotal, expenseByCat := logic.CalculateStats(*data, "expense", userId)
			ui.PrintReport(incomeTotal, incomeByCat, expenseTotal, expenseByCat)
		case "2":
			ui.ListTransactions(data.Transactions, userId)
		case "3":
			ui.AddTransaction(data, userId)
			storage.SaveDataWithCheck(data)
		case "4":
			ui.EditTransaction(data, userId)
			storage.SaveDataWithCheck(data)
		case "5":
			ui.DeleteTransaction(data, userId)
			storage.SaveDataWithCheck(data)
		case "0":
			fmt.Println("Возвращаемся в главное меню...")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова!")
		}
	}
}

func manageProfile(data *models.Database, access *bool, userId *int) {
	for {
		fmt.Println("\n=== Управление профилем ===")
		fmt.Println("1. Сменить пользователя")
		fmt.Println("2. Редактировать пользователя")
		fmt.Println("3. Удалить пользователя")
		fmt.Println("0. Назад в главное меню")
		fmt.Print("Выберите действие: ")

		profileChoice := ui.Input()

		switch profileChoice {
		case "1":
			newAccess, newUserId := ui.Authefication(data)
			storage.SaveDataWithCheck(data)
			if newUserId == -1 {
				return
			}
			if newAccess {
				*access = newAccess
				*userId = newUserId
			}
		case "2":
			ui.EditProfile(data, *userId)
			storage.SaveDataWithCheck(data)
		case "3":
			deleted := ui.DeleteUser(data, *userId)
			if deleted {
				newAccess, newUserId := ui.Authefication(data)
				storage.SaveDataWithCheck(data)
				if newUserId == -1 || newUserId == 0 {
					return
				}
				*access = newAccess
				*userId = newUserId
			}
		case "0":
			fmt.Println("Возвращаемся в главное меню...")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова!")
		}
	}
}

func handleExit(data *models.Database) {
	fmt.Println("Вы уверены, что хотите выйти? (y/n)")
	confirm := ui.Input()
	if confirm == "y" || confirm == "Y" {
		fmt.Println("Сохранение данных и выход...")
		if err := storage.SaveData("data.json", *data); err != nil {
			fmt.Printf("КРИТИЧЕСКАЯ ОШИБКА: не удалось сохранить данные: %v\n", err)
		}
	} else {
		fmt.Println("Отмена выхода. Возвращаемся в меню...")
	}
}

package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testfile/internal/helpers"
	"testfile/internal/logic"
	"testfile/internal/models"
	"testfile/internal/security"
	"time"
)

func Input() string {
	reader := bufio.NewReader(os.Stdin)
	inputValue, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	return strings.TrimSpace(inputValue)
}

func InputDate() string {
	for {
		fmt.Println("Когда операция была произведена?")
		fmt.Println("Запишите дату в формате (YYYY-MM-DD)")
		transDate := Input()
		if transDate == "0" {
			return "0"
		}
		_, err := time.Parse("2006-01-02", transDate)
		if err == nil {
			return transDate
		}
		fmt.Printf("Ошибка: дата '%s' некорректна. Убедитесь, что:\n", transDate)
		fmt.Println("- Формат: ГГГГ-ММ-ДД (например, 2024-12-31)")
		fmt.Println("- Месяц от 01 до 12")
		fmt.Println("- День соответствует месяцу (в феврале 28/29 дней)")
	}
}

func InputAmount() int {
	for {
		fmt.Println("Какова сумма данной операции")
		amountStr := Input()
		amount, err := strconv.Atoi(amountStr)
		if err == nil && amount > 0 {
			return amount
		}
		fmt.Println("Ошибка! Введите положительное число.")
	}
}

func InputCat(transactionType string, d *models.Database) string {
	for {
		if transactionType == "income" {
			fmt.Println("Укажите, к какой категории дохода относится Ваша операция")

			// Используем категории из настроек
			categories := d.Settings.Categories["income"]
			for i, cat := range categories {
				fmt.Printf("%d. %s\n", i+1, cat.Name)
			}
			fmt.Println("0. Назад")

			choice := Input()
			if CheckCancel(choice, "Отменяю выбор категории...") {
				return "0"
			}

			idx, err := strconv.Atoi(choice)
			if err != nil || idx < 1 || idx > len(categories) {
				fmt.Println("Ошибка: укажите существующую категорию дохода")
				continue
			}
			return categories[idx-1].ID
		} else {
			fmt.Println("Укажите, к какой категории расхода относится Ваша операция")

			categories := d.Settings.Categories["expense"]
			for i, cat := range categories {
				fmt.Printf("%d. %s\n", i+1, cat.Name)
			}
			fmt.Println("0. Назад")

			choice := Input()
			if CheckCancel(choice, "Отменяю выбор категории...") {
				return "0"
			}

			idx, err := strconv.Atoi(choice)
			if err != nil || idx < 1 || idx > len(categories) {
				fmt.Println("Ошибка: укажите существующую категорию расхода")
				continue
			}
			return categories[idx-1].ID
		}
	}
}

func InputType() string {
	for {
		fmt.Println("Какой тип транзакции хотите добавить? (income/expense)")
		transactionType := Input()
		if transactionType == "0" {
			return "0"
		}
		if transactionType == "income" || transactionType == "expense" {
			return transactionType
		}
		if transactionType == "доход" || transactionType == "Доход" {
			return "income"
		}
		if transactionType == "расход" || transactionType == "Расход" {
			return "expense"
		}
		fmt.Println("Ошибка: укажите существующий тип транзакции")
	}
}

func AddTransaction(d *models.Database, userId int) {
	var newTransaction models.Transaction
	fmt.Println("Чтобы отменить добавление и вернуться назад, введите '0' в любом поле")

	newTransaction.Id = logic.GetNextID(d.Transactions, userId)
	newTransaction.UserID = userId

	newTransaction.Date = InputDate()
	if CheckCancel(newTransaction.Date, "Отменяю добавление транзакции...") {
		return
	}

	newTransaction.Type = InputType()
	if CheckCancel(newTransaction.Type, "Отменяю добавление транзакции...") {
		return
	}

	newTransaction.Category = InputCat(newTransaction.Type, d)
	if CheckCancel(newTransaction.Category, "Отменяю добавление транзакции...") {
		return
	}

	fmt.Println("Расскажите подробнее, какое событие хотите добавить")
	fmt.Println("(или введите '0' для отмены)")
	newTransaction.Description = Input()
	if CheckCancel(newTransaction.Description, "Отменяю добавление транзакции...") {
		return
	}

	fmt.Println("Какова сумма данной операции? (введите 0 для отмены)")
	amountStr := Input()
	if amountStr == "0" {
		fmt.Println("Отменяю добавление транзакции...")
		return
	}
	amount, err := strconv.Atoi(amountStr)
	for err != nil || amount <= 0 {
		fmt.Println("Ошибка! Введите положительное число или 0 для отмены:")
		amountStr = Input()
		if amountStr == "0" {
			fmt.Println("Отменяю добавление транзакции...")
			return
		}
		amount, err = strconv.Atoi(amountStr)
	}
	newTransaction.Amount = amount

	d.Transactions = append(d.Transactions, newTransaction)
	userName := helpers.GetUserName(*d, userId)
	fmt.Printf("Поздравляю, %v, Вы успешно добавили новую транзакцию!\n", userName)
	fmt.Println("Более детальная информация:")
	fmt.Printf("Id: %v\nДата совершения: %v\nТип транзакции: %v\nСобытие: %v\nКатегория: %v\nСумма: %v₽\n",
		newTransaction.Id, newTransaction.Date, newTransaction.Type,
		newTransaction.Description, newTransaction.Category, newTransaction.Amount)
}
func ListTransactions(transactions []models.Transaction, userId int) {
	transactionCount := false
	for _, v := range transactions {
		if userId == v.UserID {
			transactionCount = true
		}
	}
	if !transactionCount {
		fmt.Println("Транзакций нет.")
		return
	}
	fmt.Printf("\n%-4s %-12s %-10s %-25s %-20s %10s\n",
		"ID", "Дата", "Тип", "Описание", "Категория", "Сумма")
	for _, t := range transactions {
		if t.UserID == userId {
			typeStr := "Доход"
			if t.Type == "expense" {
				typeStr = "Расход"
			}
			catName := models.CategoryNames[t.Category]
			if catName == "" {
				catName = t.Category
			}
			fmt.Printf("%-4d %-12s %-10s %-25s %-20s %10d₽\n",
				t.Id, t.Date, typeStr, t.Description, catName, t.Amount)
		}
	}
}

func DeleteTransaction(d *models.Database, userId int) {
	fmt.Println("1. Очистить весь список транзакций")
	fmt.Println("2. Удалить транзакцию по ID")
	fmt.Println("\n0. Вернуться назад")
	choice := Input()
	switch choice {
	case "1":
		fmt.Println("Вы уверены, что хотите удалить все транзакции? (y/n)")
		response := Input()
		if response == "y" || response == "Y" {
			remainingTransactions := make([]models.Transaction, 0, len(d.Transactions))

			for _, transaction := range d.Transactions {
				if transaction.UserID != userId {
					remainingTransactions = append(remainingTransactions, transaction)
				}
			}

			d.Transactions = remainingTransactions
			fmt.Println("Все транзакции удалены")
		} else {
			fmt.Println("Возвращаемся...")
			return
		}
	case "2":
		fmt.Print("Введите ID транзакции для удаления: \n")
		idStr := Input()
		if idStr == "0" {
			fmt.Println("Отменяю...")
			fmt.Println("Возвращаемся назад...")
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Неверный ID")
			return
		}

		for i, t := range d.Transactions {
			if t.UserID == userId {
				if t.Id == id {
					fmt.Printf("Удалить транзакцию '%s' на сумму %d₽? (y/n)", t.Description, t.Amount)
					confirm := Input()
					if confirm == "y" || confirm == "Y" {
						logic.RemoveTransaction(d, i)
						fmt.Println("Вы успешно удалили транзакцию")
						return
					}
				}
			}
		}
		fmt.Println("Транзакция с таким ID не найдена.")
	case "0":
		return
	default:
		fmt.Println("Ошибка выбора, попробуйте еще раз")
	}

}

func PrintReport(incomeTotal int, incomeByCat map[string]int, expenseTotal int, expenseByCat map[string]int) {
	fmt.Println("📊 Отчёт")
	fmt.Printf("Всего заработано:   +%d₽\n", incomeTotal)
	fmt.Printf("Всего потрачено:    -%d₽\n", expenseTotal)
	fmt.Printf("Итоговый баланс:    %+d₽\n\n", incomeTotal-expenseTotal)

	fmt.Println("1. Расчитать отчёт по категориям?\n2. Вернуться назад?")
	response := Input()
	if response == "1" {
		fmt.Println("ДОХОДЫ по категориям:")
		if len(incomeByCat) == 0 {
			fmt.Println("Транзакций нет")
		}
		for cat, amount := range incomeByCat {
			name := models.CategoryNames[cat]
			if name == "" {
				name = cat
			}
			fmt.Printf("  %-20s +%d₽\n", name+":", amount)
		}

		fmt.Println("\nРАСХОДЫ по категориям:")
		if len(expenseByCat) == 0 {
			fmt.Println("Транзакций нет")
		}
		for cat, amount := range expenseByCat {
			name := models.CategoryNames[cat]
			if name == "" {
				name = cat
			}
			fmt.Printf("  %-20s -%d₽\n", name+":", amount)
		}
	}
}
func CheckCancel(value string, message string) bool {
	if value == "0" {
		fmt.Println(message)
		return true
	}
	return false
}
func EditTransaction(d *models.Database, userId int) {
	transLen := 0
	fmt.Println("Какую транзакцию будем редактировать?")
	for _, v := range d.Transactions {
		if userId == v.UserID {
			transLen++
			fmt.Printf("%v: %v -> %v₽ \n", v.Id, v.Description, v.Amount)
		}
	}
	if transLen == 0 {
		fmt.Println("Транзакции не найдены")
		return
	}
	choiceStr := Input()
	choice1, err := strconv.Atoi(choiceStr)
	if err != nil {
		log.Printf("Ошибка: %v", err)
	}
	idx := -1
	for i, t := range d.Transactions {
		if userId == t.UserID {
			if t.Id == choice1 {
				idx = i
				break
			}
		}
	}

	if idx == -1 {
		fmt.Println("Транзакция не найдена")
		return
	}
	for {
		fmt.Println("Открываю...")
		fmt.Println("Какой параметр хотите изменить?")
		fmt.Println("1. Дата")
		fmt.Println("2. Тип")
		fmt.Println("3. Категория")
		fmt.Println("4. Описание")
		fmt.Println("5. Сумма")
		fmt.Println("\n0. Вернуться назад")
		choice2 := Input()
		switch choice2 {
		case "1":
			d.Transactions[idx].Date = InputDate()
		case "2":
			d.Transactions[idx].Type = InputType()
		case "3":
			d.Transactions[idx].Category = InputCat(d.Transactions[idx].Type, d)
		case "4":
			fmt.Println("Напишите новое описание транзакции")
			d.Transactions[idx].Description = Input()
		case "5":
			d.Transactions[idx].Amount = InputAmount()
		case "0":
			fmt.Println("Возвращаемся в меню...")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте еще раз!")
		}
		fmt.Println("Изменения проведены упешно!")
		fmt.Println("Продолжим изменения? (y/n)")

		confirm := Input()
		if confirm == "y" || confirm == "Y" {
			fmt.Println("Вернемся к изменению Вашей транзакции")
		} else {
			fmt.Println("Операция завершена!")
			fmt.Println("Возвращаемся в главное меню")
			return
		}
	}

}
func Authefication(d *models.Database) (bool, int) {
	fmt.Println("Личный финансовый менеджер")
	fmt.Println("Чтобы выйти из программы, введите '0'")
	fmt.Println("1. Войти")
	fmt.Println("2. Зарегестрироваться")
	choice := Input()
	switch choice {
	case "1":
		attempts := 0
		maxAttempts := 3
		fmt.Print("Введите логин пользователя: ")
		for {
			login := Input()
			if login == "0" {
				return false, -1
			}
			for _, v := range d.Users {
				if v.Login == login {
					fmt.Printf("Здравствуйте, %v\n", v.Name)

					for attempts < maxAttempts {
						fmt.Print("Введите свой пароль: ")
						password := Input()
						if password == "0" {
							return false, -1
						}

						if security.CheckPassword(password, v.Password) {
							fmt.Printf("\n%v, добро пожаловать!\n", v.Name)
							return true, v.ID
						} else {
							attempts++
							if attempts < maxAttempts {
								fmt.Printf("Неверный пароль. Осталось попыток: %d\n", maxAttempts-attempts)
							}
						}
					}

					fmt.Println("Превышено количество попыток. Попробуйте позже.")
					return false, 0
				}
			}
			fmt.Println("Неверный логин, попробуйте еще раз")
		}
	case "2":
		return true, Registration(d)
	case "0":
		return false, -1
	default:
		fmt.Println("Нет такого варианта! Попробуйте еще разы")
		return false, 0
	}

}
func Registration(d *models.Database) int {
	var newUser models.User
	fmt.Printf("Чтобы выйти нажмите '0'\n\n")
	fmt.Println("Как Вас зовут?")
	newUser.Name = Input()
	if newUser.Name == "0" {
		return -1
	}
	fmt.Println("Придумайте логин")
	for {
		login := Input()
		if login == "0" {
			return -1
		}
		if logic.ValidLogin(login, *d) != "" {
			newUser.Login = login
			break
		}
	}
	for {
		fmt.Println("Придумайте пароль (минимум 8 символов, заглавная буква, цифра, спецсимвол): ")
		password := Input()
		if password == "0" {
			return -1
		}
		if logic.ValidatePassword(password) == nil {
			newUser.Password = security.HashPassword(password)
			break
		}
	}
	newUser.ID = logic.GetNextUserID(*d)
	d.Users = append(d.Users, newUser)
	fmt.Printf("Добро пожаловать, %v", newUser.Name)
	return newUser.ID
}
func DeleteUser(d *models.Database, userId int) bool {
	fmt.Println("Вы точно хотите безвозвратно удалить свой аккаунт?")
	choice := Input()
	if choice == "y" || choice == "Y" {
		var deletedUser []models.User
		for _, v := range d.Users {
			if userId != v.ID {
				deletedUser = append(deletedUser, v)
			}
		}
		d.Users = deletedUser

		var remainingTransactions []models.Transaction
		for _, t := range d.Transactions {
			if t.UserID != userId {
				remainingTransactions = append(remainingTransactions, t)
			}
		}
		d.Transactions = remainingTransactions

		fmt.Println("Пользователь и все его транзакции удалены!")
		return true
	} else {
		fmt.Println("Отменяю удаление!")
		return false
	}
}
func EditProfile(d *models.Database, userId int) {

	userIdx := -1
	for i, u := range d.Users {
		if u.ID == userId {
			userIdx = i
			break
		}
	}

	if userIdx == -1 {
		fmt.Println("Пользователь не найден!")
		return
	}

	for {
		fmt.Printf("Что хотите изменить?\n\n")
		fmt.Println("1. Имя")
		fmt.Println("2. Логин")
		fmt.Println("3. Пароль")
		fmt.Println("\n0. Вернуться назад")
		choice := Input()
		switch choice {
		case "1":
			fmt.Printf("Ваше прошлое имя - %v\n", helpers.GetUserName(*d, userId))
			fmt.Printf("Введите новое имя: ")
			d.Users[userIdx].Name = Input()
			fmt.Printf("Ваше имя - %v\n", d.Users[userIdx].Name)
		case "2":
			for {
				fmt.Printf("Ваш прошлый логин - %v\n", d.Users[userIdx].Login)
				fmt.Printf("Введите новый логин: ")
				newLogin := Input()
				if logic.ValidLogin(newLogin, *d) != "" {
					d.Users[userIdx].Login = newLogin
					break
				}
			}
			fmt.Printf("Ваш логин - %v\n", d.Users[userIdx].Login)
		case "3":
			for {
				fmt.Println("Введите новый пароль: ")
				newPassword := Input()
				if logic.ValidatePassword(newPassword) == nil {
					d.Users[userIdx].Password = security.HashPassword(newPassword)
					break
				}
			}
		case "0":
			fmt.Println("Возвращаемся назад...")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте еще раз!")
		}
		fmt.Println("Изменения проведены упешно!")
		fmt.Println("Продолжим изменения? (y/n)")

		confirm := Input()
		if confirm == "y" || confirm == "Y" {
			fmt.Println("Вернемся к изменению Вашего профиля")
		} else {
			fmt.Println("Операция завершена!")
			fmt.Println("Возвращаемся в главное меню")
			return
		}
	}

}

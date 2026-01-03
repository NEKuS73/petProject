package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"tetsfile/internal/logic"
	"tetsfile/internal/models"
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

func InputCat(transactionType string) string {
	for {
		if transactionType == "income" {
			fmt.Println("Укажите, к какой категории дохода относится Ваша операция")
			fmt.Println("1. Зарплата")
			fmt.Println("2. Фриланс")
			fmt.Println("3. Инвестиции")
			fmt.Println("4. Гос. выплаты")
			fmt.Println("5. Прочие доходы")
			fmt.Println("0. Назад")
			category := Input()
			switch category {
			case "1":
				return "salary"
			case "2":
				return "freelance"
			case "3":
				return "investments"
			case "4":
				return "state"
			case "5":
				return "other_income"
			case "0":
				return "0"
			}
			fmt.Println("Ошибка: укажите существующую категорию дохода")
		} else {
			fmt.Println("Укажите, к какой категории расхода относится Ваша операция")
			fmt.Println("1. Жильё")
			fmt.Println("2. Еда")
			fmt.Println("3. Транспорт")
			fmt.Println("4. Здоровье")
			fmt.Println("5. Образование")
			fmt.Println("6. Обязательные платежи")
			fmt.Println("7. Прочие расходы")
			fmt.Println("0. Назад")
			category := Input()
			switch category {
			case "1":
				return "housing"
			case "2":
				return "food"
			case "3":
				return "transport"
			case "4":
				return "health"
			case "5":
				return "education"
			case "6":
				return "obligatory"
			case "7":
				return "other"
			case "0":
				return "0"
			}
			fmt.Println("Ошибка: укажите существующую категорию расхода")
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

func AddTransaction(d *models.Data) {
	var newTransaction models.Transaction
	fmt.Println("Чтобы отменить добавление и вернуться назад, Напишите '0'")

	newTransaction.Id = logic.GetNextID(d.Transactions)

	newTransaction.Date = InputDate()
	if CheckCancel(newTransaction.Date, "Отменяю...") {
		return
	}

	newTransaction.Type = InputType()
	if CheckCancel(newTransaction.Type, "Отменяю...") {
		return
	}

	newTransaction.Category = InputCat(newTransaction.Type)
	if CheckCancel(newTransaction.Category, "Отменяю...") {
		return
	}

	fmt.Println("Расскажите подробнее, какое событие хотите добавить")
	newTransaction.Description = Input()
	if CheckCancel(newTransaction.Description, "Отменяю...") {
		return
	}

	newTransaction.Amount = InputAmount()

	d.Transactions = append(d.Transactions, newTransaction)
	fmt.Printf("Поздравляю, %v, Вы успешно добавили новую транзакцию!\n", d.User)
	fmt.Println("Более детальная информация:")
	fmt.Printf("Id: %v\nДата совершения: %v\nТип транзакции: %v\nСобытие: %v\nКатегория: %v\nСумма: %v₽\n",
		newTransaction.Id, newTransaction.Date, newTransaction.Type,
		newTransaction.Description, newTransaction.Category, newTransaction.Amount)
}
func ListTransactions(transactions []models.Transaction) {
	if len(transactions) == 0 {
		fmt.Println("Транзакций нет.")
		return
	}
	fmt.Printf("\n%-4s %-12s %-10s %-25s %-20s %10s\n",
		"ID", "Дата", "Тип", "Описание", "Категория", "Сумма")
	for _, t := range transactions {
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

func DeleteTransaction(d *models.Data) {
	fmt.Println("1. Очистить весь список транзакций")
	fmt.Println("2. Удалить транзакцию по ID")
	fmt.Println("\n0. Вернуться назад")
	choice := Input()
	switch choice {
	case "1":
		fmt.Printf("Вы уверены, что хотите удалить все транзакции за %v? (y/n)\n", d.Month)
		response := Input()
		if response == "y" || response == "Y" {
			d.Transactions = nil
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
			if t.Id == id {
				fmt.Printf("Удалить транзакцию '%s' на сумму %d₽? (y/n)", t.Description, t.Amount)
				confirm := Input()
				if confirm == "y" || confirm == "Y" {
					removeTransaction(d, i)
					fmt.Println("Вы успешно удалили транзакцию")
					return
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
func removeTransaction(d *models.Data, i int) {
	d.Transactions = append(d.Transactions[:i], d.Transactions[i+1:]...)
}

func PrintReport(month string, incomeTotal int, incomeByCat map[string]int, expenseTotal int, expenseByCat map[string]int) {
	fmt.Printf("📊 Отчёт за %s\n\n", month)
	fmt.Printf("Всего заработано:   +%d₽\n", incomeTotal)
	fmt.Printf("Всего потрачено:    -%d₽\n", expenseTotal)
	fmt.Printf("Итоговый баланс:    %+d₽\n\n", incomeTotal-expenseTotal)

	fmt.Println("1. Расчитать отчёт по категориям?\n2. Вернуться назад?")
	response := Input()
	if response == "1" {
		fmt.Println("ДОХОДЫ по категориям:")
		for cat, amount := range incomeByCat {
			name := models.CategoryNames[cat]
			if name == "" {
				name = cat
			}
			fmt.Printf("  %-20s +%d₽\n", name+":", amount)
		}

		fmt.Println("\nРАСХОДЫ по категориям:")
		for cat, amount := range expenseByCat {
			name := models.CategoryNames[cat]
			if name == "" {
				name = cat
			}
			fmt.Printf("  %-20s -%d₽\n", name+":", amount)
		}
	}
}
func CheckCancel(value interface{}, message string) bool {
	switch v := value.(type) {
	case string:
		if v == "0" {
			fmt.Println(message)
			fmt.Println("Возвращаемся назад...")
			return true
		}
	case int:
		if v == 0 {
			fmt.Println(message)
			fmt.Println("Возвращаемся назад...")
			return true
		}
	case float64:
		if v == 0 {
			fmt.Println(message)
			fmt.Println("Возвращаемся назад...")
			return true
		}
	}
	return false
}
func EditTransaction(d *models.Data) {
	fmt.Println("Какую транзакцию будем редактировать?")
	for _, v := range d.Transactions {
		fmt.Printf("%v: %v -> %v₽ \n", v.Id, v.Description, v.Amount)
	}
	choiceStr := Input()
	choice1, err := strconv.Atoi(choiceStr)
	if err != nil {
		log.Printf("Ошибка: %v", err)
	}
	idx := -1
	for i, t := range d.Transactions {
		if t.Id == choice1 {
			idx = i
			break
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
			d.Transactions[idx].Category = InputCat(d.Transactions[idx].Type)
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

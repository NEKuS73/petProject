package helpers

import (
	"fmt"
	"strings"
	"tetsfile/internal/models"
)

func GetUserName(d models.Database, userId int) string {
	userName := "Гость"
	for _, v := range d.Users {
		if v.ID == userId {
			userName = v.Name
		}
	}
	return userName
}
func GetDate(dateStr string) string {
	listOfDate := strings.Split(dateStr, "-")
	if len(listOfDate) < 3 {
		return dateStr
	}

	year := listOfDate[0]
	month := listOfDate[1]
	day := listOfDate[2]

	var monthName string
	switch month {
	case "01":
		monthName = "Январь"
	case "02":
		monthName = "Февраль"
	case "03":
		monthName = "Март"
	case "04":
		monthName = "Апрель"
	case "05":
		monthName = "Май"
	case "06":
		monthName = "Июнь"
	case "07":
		monthName = "Июль"
	case "08":
		monthName = "Август"
	case "09":
		monthName = "Сентябрь"
	case "10":
		monthName = "Октябрь"
	case "11":
		monthName = "Ноябрь"
	case "12":
		monthName = "Декабрь"
	default:
		monthName = month
	}

	return fmt.Sprintf("%s %s, %s", day, monthName, year)
}
func FilterTransactionsByMonth(transactions []models.Transaction, yearMonth string) []models.Transaction {
	var result []models.Transaction
	for _, t := range transactions {
		if strings.HasPrefix(t.Date, yearMonth) {
			result = append(result, t)
		}
	}
	return result
}

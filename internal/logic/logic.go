package logic

import (
	"errors"
	"fmt"
	"strings"
	"tetsfile/internal/models"
)

func CalculateStats(d models.Database, transType string, userId int) (int, map[string]int) {
	total := 0
	byCategory := make(map[string]int)
	for _, t := range d.Transactions {
		if t.Type == transType && t.UserID == userId {
			total += t.Amount
			byCategory[t.Category] += t.Amount
		}
	}
	return total, byCategory
}

func GetNextID(transactions []models.Transaction, userId int) int {
	maxID := 0
	for _, t := range transactions {
		if userId == t.UserID {
			if t.Id > maxID {
				maxID = t.Id
			}
		}
	}
	return maxID + 1
}
func GetNextUserID(d models.Database) int {
	maxID := 0
	for _, v := range d.Users {
		if v.ID >= maxID {
			maxID = v.ID
		}
	}
	return maxID + 1
}
func RemoveTransaction(d *models.Database, i int) {
	d.Transactions = append(d.Transactions[:i], d.Transactions[i+1:]...)
}
func ValidLogin(login string, d models.Database) string {

	for _, v := range d.Users {
		if login == v.Login {
			fmt.Println("Такой логин уже занят, придумайте другой")
			return ""
		}
	}
	return login
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		fmt.Println("Пароль должен содержать не менее 8 символов")
		return errors.New("")
	}

	hasUpper := false
	hasDigit := false
	hasSpecial := false

	specialChars := "`~!@#$%^&*()_-+={}[]\\|:;'<>,.?/"

	for _, char := range password {
		if 'A' <= char && char <= 'Z' {
			hasUpper = true
		}
		if '0' <= char && char <= '9' {
			hasDigit = true
		}
		if strings.ContainsRune(specialChars, char) {
			hasSpecial = true
		}

		if hasUpper && hasDigit && hasSpecial {
			break
		}
	}

	if !hasUpper {
		fmt.Println("Пароль должен содержать хотя бы одну заглавную букву")
		return errors.New("")
	}
	if !hasDigit {
		fmt.Println("Пароль должен содержать хотя бы одну цифру (0-9)")
		return errors.New("")
	}
	if !hasSpecial {
		fmt.Println("Пароль должен содержать хотя бы один спецсимвол")
		return errors.New("")
	}

	return nil
}

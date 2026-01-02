package logic

import "tetsfile/internal/models"

func CalculateStats(d models.Data, transType string) (int, map[string]int) {
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

func GetNextID(transactions []models.Transaction) int {
	maxID := 0
	for _, t := range transactions {
		if t.Id > maxID {
			maxID = t.Id
		}
	}
	return maxID + 1
}

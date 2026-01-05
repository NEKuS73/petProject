package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"tetsfile/internal/models"
)

func SaveData(fileName string, d models.Database) error {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка преобразования в JSON: %w", err)
	}

	if err := os.WriteFile(fileName, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}

	return nil
}
func SaveDataWithCheck(data *models.Database) {
	if err := SaveData("data.json", *data); err != nil {
		fmt.Printf("КРИТИЧЕСКАЯ ОШИБКА: не удалось сохранить данные: %v\n", err)
	}
}

func LoadData(fileName string) models.Database {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Внимание: файл %s не найден. Создаю новый.\n", fileName)
		return CreateNewDB()
	}

	var data models.Database
	if err := json.Unmarshal(file, &data); err != nil {
		fmt.Printf("ERROR: некорректный формат файла: %v\n", err)
		return CreateNewDB()
	}
	return data
}
func CreateNewDB() models.Database {
	return models.Database{
		Users: []models.User{},

		Settings: models.Settings{
			DefaultCurrency: "RUB",
			Categories: map[string][]models.Category{
				"income": {
					{ID: "salary", Name: "Зарплата"},
					{ID: "freelance", Name: "Фриланс"},
					{ID: "other_income", Name: "Прочие доходы"},
				},
				"expense": {
					{ID: "food", Name: "Еда"},
					{ID: "transport", Name: "Транспорт"},
					{ID: "other", Name: "Прочие расходы"},
				},
			},
		},
	}
}

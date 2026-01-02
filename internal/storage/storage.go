package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"tetsfile/internal/models"
)

func SaveData(fileName string, d models.Data) error {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка преобразования в JSON: %w", err)
	}

	if err := os.WriteFile(fileName, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}

	return nil
}

func LoadData(fileName string) models.Data {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Внимание: файл %s не найден. Создаю новый.\n", fileName)
		return models.Data{
			User:  "Гость",
			Month: "Текущий месяц",
		}
	}

	var data models.Data
	if err := json.Unmarshal(file, &data); err != nil {
		fmt.Printf("ERROR: некорректный формат файла: %v\n", err)
		return models.Data{
			User:  "Гость",
			Month: "Текущий месяц",
		}
	}
	return data
}

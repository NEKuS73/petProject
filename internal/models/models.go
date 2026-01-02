package models

type Transaction struct {
	Id          int    `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
	Type        string `json:"type"`
}

type Data struct {
	User         string        `json:"user"`
	Month        string        `json:"month"`
	Transactions []Transaction `json:"transactions"`
}

var CategoryNames = map[string]string{
	"salary":       "Зарплата",
	"freelance":    "Фриланс",
	"investments":  "Инвестиции",
	"state":        "Гос. выплаты",
	"other_income": "Прочие доходы",
	"housing":      "Жильё",
	"food":         "Еда",
	"transport":    "Транспорт",
	"health":       "Здоровье",
	"education":    "Образование",
	"obligatory":   "Обязательные платежи",
	"other":        "Прочие расходы",
}

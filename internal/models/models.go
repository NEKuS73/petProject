package models

type Transaction struct {
	Id          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
	Type        string `json:"type"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Database struct {
	Users        []User        `json:"users"`
	Transactions []Transaction `json:"transactions"`
	Settings     Settings      `json:"settings"`
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

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Settings struct {
	DefaultCurrency string                `json:"default_currency,omitempty"`
	Categories      map[string][]Category `json:"categories"`
}

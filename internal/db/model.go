package db

type Card struct {
	Number string `json:"number"`
	Date   string `json:"date"`
	CVV    string `json:"CVV"`
}

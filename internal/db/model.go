package db

type Card struct {
	Number string `json:"number"`
	Date   string `json:"date"`
	CVV    string `json:"CVV"`
}

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

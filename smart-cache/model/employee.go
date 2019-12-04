package model

type Person struct {
	ID       string `json:"_id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	City     string `json:"city"`
}

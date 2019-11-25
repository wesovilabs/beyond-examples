package model

type Employee struct {
	ID       string `json:"_id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
}

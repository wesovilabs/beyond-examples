package model

type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"-"`
}


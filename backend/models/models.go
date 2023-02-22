package models

import "time"

type GroceryPageProduct struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

type GroceryListSelections struct {
	Selections []int `json:"selections"`
}

type Account struct {
	ID        int       `json:"id"`
	UserName  string    `json:"uname"`
	Epw       string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type LoginRequest struct {
	Username string `json:"uname"`
	Password string `json:"pw"`
}

type LoginResponse struct{}

type CreateAccountRequest struct{}

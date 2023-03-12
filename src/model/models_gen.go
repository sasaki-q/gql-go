// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Item struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Price int    `json:"price"`
	User  *User  `json:"user"`
}

type NewItem struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
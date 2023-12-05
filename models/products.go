package models

import (
	"fmt"
	"net/http"
)

type Product struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
}

type ProductsList struct {
	Products []Product `json:"products"`
}

func (i *Product) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")

	}
	return nil
}

func (i *Product) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (i *ProductsList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

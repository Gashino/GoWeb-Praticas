package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func LoadData() ([]Product, error) {
	var products []Product

	file, err := os.ReadFile("/Users/agaggino/Documents/Bootcamp/GoWeb/products.json")

	if err != nil {
		return nil, errors.New("error: an error occurred with the file...")
	}

	errJson := json.Unmarshal(file, &products)

	if errJson != nil {
		return nil, errors.New("error: an erorr ocurred while parsing the file...")
	}

	return products, nil
}

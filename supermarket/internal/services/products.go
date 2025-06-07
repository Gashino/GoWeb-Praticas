package services

import (
	"encoding/json"
	"errors"
	"local/go-web/supermarket/internal/domain/entities"
	"os"
)

type Interface interface {
}

type ProductInterface interface {
	Interface
	GetAll(d []entities.Product) ([]entities.Product, error)
	GetById(d []entities.Product, id int) (*entities.Product, error)
	GetGreaterThan(d []entities.Product, gt float64) (*[]entities.Product, error)
	PostProduct(d *[]entities.Product, p entities.Product) (bool, error)
}

var Data, _ = LoadData()

// Methods
func LoadData() ([]entities.Product, error) {
	var products []entities.Product

	file, err := os.ReadFile("products.json")

	if err != nil {
		return nil, errors.New("error: an error occurred with the file...")
	}

	errJson := json.Unmarshal(file, &products)

	if errJson != nil {
		return nil, errors.New("error: an erorr ocurred while parsing the file...")
	}

	return products, nil
}

package repository

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
	GetAll() ([]entities.Product, error)
	GetById(id int) (*entities.Product, error)
	GetGreaterThan(gt float64) (*[]entities.Product, error)
	PostProduct(p entities.Product) (bool, error)
	DeleteProduct(id int) error
	UpdateProduct(p entities.Product) error
}

func NewMemoryRepo() *ProductBd {

	data, err := loadData()

	if err != nil {
		panic(err)
	}

	return &ProductBd{
		data: &data,
	}
}

type ProductBd struct {
	data *[]entities.Product
}

func (p *ProductBd) GetAll() ([]entities.Product, error) {
	return *p.data, nil
}

func (p *ProductBd) GetById(id int) (*entities.Product, error) {
	for _, prod := range *p.data {
		if prod.Id == id {
			return &prod, nil
		}
	}
	return nil, errors.New("error: inexistent entities.Product")
}

func (p *ProductBd) GetGreaterThan(gt float64) (*[]entities.Product, error) {
	var products []entities.Product

	for _, prod := range *p.data {
		if prod.Price > gt {
			products = append(products, prod)
		}
	}
	return &products, nil
}

func (p *ProductBd) PostProduct(prod entities.Product) (bool, error) {
	idToGive := (*p.data)[len(*p.data)-1].Id
	prod.Id = idToGive + 1
	*p.data = append(*p.data, prod)
	return true, nil
}

func (p *ProductBd) DeleteProduct(id int) error {

	for i := range *p.data {
		if (*p.data)[i].Id == id {
			*p.data = append((*p.data)[:i], (*p.data)[i+1:]...)
			return nil
		}
	}

	return errors.New("error: product not found")
}

func (p *ProductBd) UpdateProduct(prod entities.Product) error {
	for i := range *p.data {
		if (*p.data)[i].Id == prod.Id {
			(*p.data)[i] = prod
			return nil
		}
	}

	return errors.New("error: product not found")
}

// Methods
func loadData() ([]entities.Product, error) {
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

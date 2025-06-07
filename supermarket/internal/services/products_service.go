package services

import (
	"errors"
	"local/go-web/supermarket/internal/domain/entities"
)

// Interface Methods
type ProductService struct {
}

func (p ProductService) PostProduct(d *[]entities.Product, pp entities.Product) (bool, error) {
	idToGive := (*d)[len(*d)-1].Id
	pp.Id = idToGive + 1
	*d = append(*d, pp)
	return true, nil
}

func (p ProductService) GetAll(d []entities.Product) ([]entities.Product, error) {
	return d, nil
}

func (p ProductService) GetById(d []entities.Product, id int) (*entities.Product, error) {
	for _, prod := range d {
		if prod.Id == id {
			return &prod, nil
		}
	}
	return nil, errors.New("error: inexistent entities.Product")
}

func (p ProductService) GetGreaterThan(d []entities.Product, gt float64) (*[]entities.Product, error) {
	var products []entities.Product

	for _, prod := range d {
		if prod.Price > gt {
			products = append(products, prod)
		}
	}
	return &products, nil
}

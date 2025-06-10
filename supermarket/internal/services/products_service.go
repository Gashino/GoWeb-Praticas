package services

import (
	"errors"
	"local/go-web/supermarket/internal/domain/entities"
	"local/go-web/supermarket/internal/repository"
	"regexp"
)

// Interface Methods
type ProductService struct {
	repoProd repository.ProductBd
}

func StartProdService() *ProductService {
	repo := *repository.NewMemoryRepo()

	return &ProductService{
		repoProd: repo,
	}
}

func (p ProductService) PostProduct(pp entities.Product) (bool, error) {
	return p.repoProd.PostProduct(pp)
}

func (p ProductService) GetAll() ([]entities.Product, error) {
	return p.repoProd.GetAll()
}

func (p ProductService) GetById(id int) (*entities.Product, error) {
	return p.repoProd.GetById(id)
}

func (p ProductService) GetGreaterThan(gt float64) (*[]entities.Product, error) {
	return p.repoProd.GetGreaterThan(gt)
}

func (p ProductService) DeleteProduct(id int) error {
	return p.repoProd.DeleteProduct(id)
}

func (p ProductService) UpdateProduct(prod entities.Product) error {

	return p.repoProd.UpdateProduct(prod)
}

func (p ProductService) ValidateProduct(product entities.Product) error {

	if product.Price == float64(0) || product.Name == "" || product.CodeValue == "" || product.Expiration == "" || product.Quantity == 0 {
		return errors.New("error: some value is empty ")
	}

	data, _ := p.repoProd.GetAll()

	for _, v := range data {
		if v.CodeValue == product.CodeValue {
			return errors.New("error: code_value already exist")
		}
	}

	dateRegex := `^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[0-2])/\d{4}$`
	match, _ := regexp.MatchString(dateRegex, product.Expiration)

	if !match {
		return errors.New("error: expiration date doesnt match")
	}

	return nil
}

package validators

import (
	"errors"
	"local/go-web/supermarket/internal/domain/entities"
	"regexp"
)

func ValidateProduct(data []entities.Product, product entities.Product) error {

	if product.Price == float64(0) || product.Name == "" || product.CodeValue == "" || product.Expiration == "" || product.Quantity == 0 {
		return errors.New("error: some value is empty ")
	}

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

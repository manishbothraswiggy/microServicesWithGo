package validation

import (
	"regexp"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0,lt=50"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"_"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"_"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", ValidateSKU)
	return validate.Struct(p)

}

func ValidateSKU(fieldLevel validator.FieldLevel) bool {
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fieldLevel.Field().String(), -1)
	return len(matches) == 1
}

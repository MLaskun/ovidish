package model

import "github.com/MLaskun/ovidish/internal/validator"

type Product struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Quantity    int32    `json:"quantity"`
	Price       float64  `json:"price"`
	Version     int32    `json:"version"`
}

func ValidateProduct(v *validator.Validator, product *Product) {
	v.Check(product.Name != "", "name", "must be provided")

	v.Check(product.Categories != nil, "categories", "must be provided")
	v.Check(len(product.Categories) >= 1, "categories",
		"must contain at least 1 category")
	v.Check(len(product.Categories) <= 5, "categories",
		"must not contain more than 5 categories")
	v.Check(validator.Unique(product.Categories), "categories",
		"must not contain duplicate values")

	v.Check(product.Quantity >= 0, "quantity", "must not be a negative number")

	v.Check(product.Price >= 0, "price", "must not be a negative number")
}

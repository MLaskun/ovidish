package product

import (
	"fmt"

	"github.com/MLaskun/ovidish/internal/product/model"
)

type ProductService struct {
	repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetById(id int64) (*model.Product, error) {
	productModel, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	product := mapFromModel(productModel)

	return product, err
}

func (s *ProductService) Create(product *model.Product) error {
	productModel := mapToModel(product)

	err := s.repo.Insert(productModel)
	if err != nil {
		return err
	}

	fmt.Printf("product created: %v", productModel.ID)
	product.ID = productModel.ID
	product.Version = productModel.Version

	return nil
}

func mapToModel(product *model.Product) *ProductModel {
	return &ProductModel{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Categories:  product.Categories,
		Quantity:    product.Quantity,
		Price:       product.Price,
		Version:     product.Version,
	}
}

func mapFromModel(productModel *ProductModel) *model.Product {
	return &model.Product{
		ID:          productModel.ID,
		Name:        productModel.Name,
		Description: productModel.Description,
		Categories:  productModel.Categories,
		Quantity:    productModel.Quantity,
		Price:       productModel.Price,
		Version:     productModel.Version,
	}
}

package product

import (
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

func (s *ProductService) GetAll(name string,
	categories []string) ([]*model.Product, error) {
	productModelList, err := s.repo.GetAll(name, categories)
	if err != nil {
		return nil, err
	}

	products := make([]*model.Product, len(productModelList))
	for i, productModel := range productModelList {
		products[i] = mapFromModel(productModel)
	}

	return products, nil
}

func (s *ProductService) Create(product *model.Product) error {
	productModel := mapToModel(product)

	err := s.repo.Insert(productModel)
	if err != nil {
		return err
	}

	product.ID = productModel.ID
	product.Version = productModel.Version

	return nil
}

func (s *ProductService) Update(product *model.Product) error {
	productModel := mapToModel(product)

	err := s.repo.Update(productModel)
	if err != nil {
		return err
	}

	product.Version = productModel.Version

	return nil
}

func (s *ProductService) Delete(id int64) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

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

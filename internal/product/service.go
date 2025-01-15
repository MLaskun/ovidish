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
		repo: repo}
}

func (s *ProductService) GetById(id int64) (*model.Product, error) {
	return s.repo.Get(id)
}

func (s *ProductService) Create(product *model.Product) error {
	err := s.repo.Insert(product)
	if err != nil {
		return err
	}

	fmt.Println("product created")
	return nil
}

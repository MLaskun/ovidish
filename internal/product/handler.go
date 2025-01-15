package product

import (
	"net/http"

	"github.com/MLaskun/ovidish/internal/helpers"
	"github.com/MLaskun/ovidish/internal/product/model"
)

type ProductHandler struct {
	svc *ProductService
}

func NewProductHandler(svc *ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h ProductHandler) createProductHandler(w http.ResponseWriter,
	r *http.Request) {
	var input struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Categories  []string `json:"categories"`
		Quantity    int32    `json:"quantity"`
		Price       float64  `json:"price"`
	}

	err := helpers.ReadJSON(w, r, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := &model.Product{
		Name:        input.Name,
		Description: input.Description,
		Categories:  input.Categories,
		Quantity:    input.Quantity,
		Price:       input.Price,
	}

	err = h.svc.Create(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

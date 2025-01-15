package product

import (
	"net/http"

	"github.com/MLaskun/ovidish/internal/helpers"
	"github.com/MLaskun/ovidish/internal/product/model"
)

type ProductHandler struct {
	Repo ProductRepository
}

func NewProductHandler(repo ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
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
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	product := &model.Product{
		Name:        input.Name,
		Description: input.Description,
		Categories:  input.Categories,
		Quantity:    input.Quantity,
		Price:       input.Price,
	}

	err = h.Repo.Insert(product)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

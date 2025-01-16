package product

import (
	"errors"
	"net/http"

	"github.com/MLaskun/ovidish/internal/helpers"
	"github.com/MLaskun/ovidish/internal/product/model"
	"github.com/MLaskun/ovidish/internal/validator"
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

	v := validator.New()

	if model.ValidateProduct(v, product); !v.Valid() {
		helpers.FailedValidationResponse(w, r, v.Errors)
		return
	}

	err = h.svc.Create(product)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusCreated,
		helpers.Envelope{"product": product})
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

func (h ProductHandler) readProductHandler(w http.ResponseWriter,
	r *http.Request) {
	id, err := helpers.ReadIDParam(r)
	if err != nil {
		helpers.NotFoundResponse(w, r)
		return
	}

	product, err := h.svc.GetById(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrNoRecordFound):
			helpers.NotFoundResponse(w, r)
		default:
			helpers.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"product": product})
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

func (h ProductHandler) readAllProductsHandler(w http.ResponseWriter,
	r *http.Request) {
	var input struct {
		Name       string
		Categories []string
	}

	qs := r.URL.Query()

	input.Name = helpers.ReadString(qs, "name", "")
	input.Categories = helpers.ReadCSV(qs, "categories", []string{})

	products, err := h.svc.GetAll(input.Name, input.Categories)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK,
		helpers.Envelope{"products": products})
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

func (h ProductHandler) updateProductHandler(w http.ResponseWriter,
	r *http.Request) {
	id, err := helpers.ReadIDParam(r)
	if err != nil {
		helpers.NotFoundResponse(w, r)
		return
	}

	product, err := h.svc.GetById(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrNoRecordFound):
			helpers.NotFoundResponse(w, r)
		default:
			helpers.ServerErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name        *string  `json:"name"`
		Description *string  `json:"description"`
		Categories  []string `json:"categories"`
		Quantity    *int32   `json:"quantity"`
		Price       *float64 `json:"price"`
	}

	err = helpers.ReadJSON(w, r, &input)
	if err != nil {
		helpers.BadRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.Categories != nil {
		product.Categories = input.Categories
	}
	if input.Quantity != nil {
		product.Quantity = *input.Quantity
	}
	if input.Price != nil {
		product.Price = *input.Price
	}

	v := validator.New()
	if model.ValidateProduct(v, product); !v.Valid() {
		helpers.FailedValidationResponse(w, r, v.Errors)
		return
	}

	err = h.svc.Update(product)
	if err != nil {
		switch {
		case errors.Is(err, ErrEditConflict):
			helpers.EditConflictResponse(w, r)
		}
	}

	err = helpers.WriteJSON(w, http.StatusOK,
		helpers.Envelope{"product": product})
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

func (h ProductHandler) DeleteProductHandler(w http.ResponseWriter,
	r *http.Request) {
	id, err := helpers.ReadIDParam(r)
	if err != nil {
		helpers.NotFoundResponse(w, r)
		return
	}

	err = h.svc.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrNoRecordFound):
			helpers.NotFoundResponse(w, r)
		default:
			helpers.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK,
		helpers.Envelope{"message": "movie deleted successfully"})
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

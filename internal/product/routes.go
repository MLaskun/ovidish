package product

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes(handler *ProductHandler) http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/v1/product",
		handler.createProductHandler)
	router.HandlerFunc(http.MethodGet, "/v1/product/:id",
		handler.readProductHandler)
	router.HandlerFunc(http.MethodGet, "/v1/product",
		handler.readAllProductsHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/product/:id",
		handler.updateProductHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/product/:id",
		handler.DeleteProductHandler)

	return router
}

package product

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes(handler *ProductHandler) http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/v1/product",
		handler.createProductHandler)
	return router
}

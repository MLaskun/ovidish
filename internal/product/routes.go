package product

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (handler *ProductHandler) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/v1/product",
		handler.createProductHandler)
	return router
}

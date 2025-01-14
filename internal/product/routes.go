package product

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes(pr ProductRepository) http.Handler {
	router := httprouter.New()

	return router
}

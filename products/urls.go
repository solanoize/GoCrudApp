package products

import (
	"net/http"

	"github.com/go-chi/chi"
)

func URLPattern(service *ProductService) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ProductListHandler(w, r, service)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		ProductCreateHandler(w, r, service)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		ProductDetailHandler(w, r, service)
	})

	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		ProductUpdateHandler(w, r, service)
	})

	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		ProductDeleteHandler(w, r, service)
	})

	return r
}

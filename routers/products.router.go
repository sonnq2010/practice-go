package routers

import (
	"net/http"
	"practice-go/controllers"

	"github.com/go-chi/chi/v5"
)

func ProductsRouter(c controllers.IProductsController) http.Handler {
	r := chi.NewRouter()

	r.Get("/", c.GetAll)
	r.Post("/", c.Create)
	r.Get("/{productID}", c.GetByID)
	r.Put("/{productID}", c.Update)
	r.Delete("/{productID}", c.Delete)

	return r
}

func NewProductsRouter() http.Handler {
	return ProductsRouter(controllers.NewProductsController())
}

package main

import (
	"net/http"
	"practice-go/routers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)

	// Mount routers
	r.Mount("/auth", routers.NewAuthRouter())
	r.Mount("/products", routers.NewProductsRouter())

	http.ListenAndServe(":3333", r)
}

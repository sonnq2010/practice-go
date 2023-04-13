package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type IProductsController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type ProductsController struct{}

func NewProductsController() IProductsController {
	return ProductsController{}
}

func (c ProductsController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All products"))
}

func (c ProductsController) GetByID(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productID")

	w.Write([]byte("Product: " + productID))

}

func (c ProductsController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c ProductsController) Update(w http.ResponseWriter, r *http.Request) {

}

func (c ProductsController) Delete(w http.ResponseWriter, r *http.Request) {

}

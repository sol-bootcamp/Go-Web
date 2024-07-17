package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web/internal/domain"
	"web/internal/product"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service product.ProductService
}

func NewProductHandler(service product.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	products, err := ph.service.GetAllProducts()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}

func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	product, err := ph.service.GetProductByID(id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (ph *ProductHandler) SearchProduct(w http.ResponseWriter, r *http.Request) {
	priceGtStr := r.URL.Query().Get("priceGt")
	if priceGtStr == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	products, err := ph.service.SearchProduct(priceGt)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct domain.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ph.service.CreateProduct(newProduct)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web/internal/domain"
	"web/internal/service"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (ph *ProductHandler) GetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products, err := ph.service.GetAllProducts()
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)

	}
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

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = ph.service.DeleteProduct(id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "product deleted"}`))
}

func (ph *ProductHandler) UpdateCreate(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	var newProduct domain.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ph.service.UpdateProduct(id, newProduct)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (ph *ProductHandler) Patch(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	var newProduct domain.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ph.service.PatchProduct(id, newProduct)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

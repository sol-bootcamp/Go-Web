package main

import (
	"log"
	"net/http"
	"web/cmd/server/handler"
	"web/internal/product"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// inicialzacion del repo
	prodRepo, err := product.NewProductRepository("products.json")
	if err != nil {
		log.Fatalf("Error loading products: %v", err)
	}

	// inicializacion del servicio

	prodService := product.NewProductService(prodRepo)

	// inicializamos el handler
	prodHandler := handler.NewProductHandler(prodService)

	r.Get("/products", prodHandler.GetAllProducts)

	r.Get("/products/{id}", prodHandler.GetProductByID)

	r.Get("/products/search", prodHandler.SearchProduct)

	r.Post("/products", prodHandler.CreateProduct)

	log.Println("Starting server on :2000...")
	err = http.ListenAndServe(":2000", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

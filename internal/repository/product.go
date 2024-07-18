package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"web/internal/domain"
)

// ProductRepository is the interface that provides product methods
type ProductRepository interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(id int) (domain.Product, error)
	// SearchProduct(priceGt float64) ([]domain.Product, error)
	CreateProduct(product domain.Product) error
	DeleteProduct(id int) error
	UpdateProduct(int, domain.Product) error
	PatchProduct(int, domain.Product) error
}

// productRepository is a concrete implementation of ProductRepository
type productRepository struct {
	products []domain.Product
}

// NewProductRepository creates a new ProductRepository with the necessary dependencies
func NewProductRepository(filename string) (ProductRepository, error) {
	repo := &productRepository{}
	err := repo.loadProducts(filename)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (pr *productRepository) loadProducts(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pr.products); err != nil {
		return err
	}
	return nil
}
func (pr *productRepository) GetAllProducts() ([]domain.Product, error) {
	if pr.products == nil {
		return nil, fmt.Errorf("no products found")
	}
	return pr.products, nil
}

func (pr *productRepository) GetProductByID(id int) (domain.Product, error) {
	for _, product := range pr.products {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("product not found")

}

//	func (pr *productRepository) SearchProduct(priceGt float64) ([]domain.Product, error) {
//		var filteredProducts []domain.Product
//		for _, product := range pr.products {
//			if product.Price > priceGt {
//				filteredProducts = append(filteredProducts, product)
//			}
//		}
//		return filteredProducts, nil
//	}
func (pr *productRepository) CreateProduct(product domain.Product) error {
	product.ID = pr.getNextID()
	pr.products = append(pr.products, product)
	return nil
}

func (pr *productRepository) getNextID() int {
	maxID := 0
	for _, product := range pr.products {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID + 1
}

func (pr *productRepository) DeleteProduct(id int) error {
	for i, product := range pr.products {
		if product.ID == id {
			pr.products = append(pr.products[:i], pr.products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("product not found")
}

func (pr *productRepository) UpdateProduct(id int, product domain.Product) error {
	for i, p := range pr.products {
		if p.ID == id {
			pr.products[i] = product
			return nil
		}
	}
	return fmt.Errorf("product not found")
}

func (pr *productRepository) PatchProduct(id int, product domain.Product) error {
	for i, p := range pr.products {
		if p.ID == id {
			if product.Name != "" {
				pr.products[i].Name = product.Name
			}
			if product.Price != 0 {
				pr.products[i].Price = product.Price
			}
			if product.Quantity != 0 {
				pr.products[i].Quantity = product.Quantity
			}
			return nil
		}
	}
	return fmt.Errorf("product not found")
}

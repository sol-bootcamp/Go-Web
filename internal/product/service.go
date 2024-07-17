package product

import (
	"fmt"
	"web/internal/domain"
)

// ProductService is the interface that provides product methods
type ProductService interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(id int) (domain.Product, error)
	SearchProduct(priceGt float64) ([]domain.Product, error)
	CreateProduct(product domain.Product) error
}

// productService is a concrete implementation of ProductService
type productService struct {
	repository ProductRepository
}

// NewProductService creates a new ProductService with the necessary dependencies
func NewProductService(repository ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

func (ps *productService) GetAllProducts() ([]domain.Product, error) {
	product, err := ps.repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}
	return product, nil
}

func (ps *productService) GetProductByID(id int) (domain.Product, error) {
	product, err := ps.repository.GetProductByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (ps *productService) SearchProduct(priceGt float64) ([]domain.Product, error) {
	var filteredProducts []domain.Product
	products, err := ps.repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error getting all  products: %w", err)
	}
	for _, product := range products {
		if product.Price > priceGt {
			filteredProducts = append(filteredProducts, product)
		}
	}

	return filteredProducts, nil
}

func (ps *productService) CreateProduct(product domain.Product) error {
	return ps.repository.CreateProduct(product)
}

package command

import "github.com/ynuraddi/playground/inventory/internal/domain"

type (
	CreateProduct struct {
		Name     string
		Quantity int
		Price    domain.Money
	}

	ProduceService struct {
		productRepository domain.ProductRepository
	}
)

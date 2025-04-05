package command

import "inventory/internal/domain"

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

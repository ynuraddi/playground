package domain

import "context"

type ProductRepository interface {
	Save(ctx context.Context, product Product) error
	Update(ctx context.Context, product Product) error

	FindById(ctx context.Context, id string) (Product, error)
	FindByIdForUpdate(ctx context.Context, id string) (Product, error)
}

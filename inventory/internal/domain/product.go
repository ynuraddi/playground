package domain

import (
	"github.com/google/uuid"
	"github.com/stackus/errors"
)

var (
	ErrProductAlreadyExists        = errors.Wrap(errors.ErrBadRequest, "Product already exists")
	ErrProductIsNotExist           = errors.Wrap(errors.ErrBadRequest, "product is not exist")
	ErrProductQuantityCantNegative = errors.Wrap(errors.ErrBadRequest, "product quantity cant't be negative")
)

type Product struct {
	EventProducer
	Id       string
	Name     string
	Price    Money
	Quantity int64
}

func (p Product) TotalPrice() Money {
	return p.Price.Mul(p.Quantity)
}

func (p *Product) Events() (events []Event) {
	defer func() { p.events = nil }()
	return p.events
}

func NewProduct(price Money, quantity int64) Product {
	p := Product{
		Id:       uuid.NewString(),
		Price:    price,
		Quantity: quantity,
	}
	p.AddEvent(EventProductCreated{p})

	return p
}

func (p *Product) AddQuantity(quantity int64) {
	p.Quantity += quantity
}

func (p *Product) SubQuantity(quantity int64) error {
	if quantity > p.Quantity {
		return ErrProductQuantityCantNegative
	}

	p.Quantity -= quantity
	return nil
}

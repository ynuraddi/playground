package domain

type EventProductCreated struct {
	Product Product
}

func (e EventProductCreated) EventName() string { return "inventory.ProductCreated" }

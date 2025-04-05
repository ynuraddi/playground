package domain

import "fmt"

type Money struct {
	Amount int64
}

func NewMoney(amount int64) Money {
	return Money{Amount: amount}
}

func (m Money) String() string {
	return fmt.Sprintf("%d $", m.Amount)
}

func (m Money) Add(other Money) Money {
	return Money{Amount: m.Amount + other.Amount}
}

func (m Money) Sub(other Money) Money {
	return Money{Amount: m.Amount - other.Amount}
}

func (m Money) Mul(factor int64) Money {
	return Money{Amount: m.Amount * factor}
}

func (m Money) Div(factor int64) Money {
	return Money{Amount: m.Amount / factor}
}

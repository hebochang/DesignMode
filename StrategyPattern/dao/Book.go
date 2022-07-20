package dao

import "designMode/StrategyPattern/Interface"

type Book struct {
	Price float32
}

func (b *Book) GetPrice(d Interface.IDiscount) float32 {
	return b.Price * d.Discount()
}

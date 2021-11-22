package discount

type DiscountStrategy interface {
	GetValue() int
	SetValue(value int)
	GetDiscount(total float32) float32
}

type DiscountImpl struct {
	Value int
}

func (discount DiscountImpl) GetValue() int {
	return discount.Value
}

func (discount *DiscountImpl) SetValue(value int) {
	discount.Value = value
}

func (discount DiscountImpl) GetDiscount(total float32) float32 {
	return total - total * (float32(discount.Value) / 100)
}





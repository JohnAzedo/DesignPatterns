package discount

type DiscountBlackFriday struct {
	DiscountStrategy
}

func NewDiscountBlackFriday() DiscountStrategy {
	return &DiscountBlackFriday{
		&DiscountImpl{
			0,
		},
	}
}

func (discount DiscountBlackFriday) GetDiscount(total float32) float32 {
	if total >= 100 && total < 200 {
		discount.SetValue(10)
	} else if total >= 200 {
		discount.SetValue(15)
	}

	return discount.DiscountStrategy.GetDiscount(total)
}

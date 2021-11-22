package discount

type DiscountNormal struct {
	DiscountStrategy
}

func NewDiscountNormal() DiscountStrategy {
	return &DiscountNormal{
		&DiscountImpl{
			0,
		},
	}
}

func (discount DiscountNormal) GetDiscount(total float32) float32 {
	if total >= 100 && total < 200{
		discount.SetValue(5)
	} else if total >= 200 {
		discount.SetValue(10)
	}

	return discount.DiscountStrategy.GetDiscount(total)
}

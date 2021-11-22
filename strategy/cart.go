package strategy

import "DesignPatterns/strategy/discount"

type ShoppingCart interface {
	AddProduct(p *Product)
	GetTotal() float32
	GetTotalWithDiscount(strategy discount.DiscountStrategy) float32
}

type ShoppingCartImpl struct {
	Products []Product
}

func NewShoppingCart() ShoppingCart{
	return &ShoppingCartImpl{}
}

func (purchase *ShoppingCartImpl) AddProduct(p *Product) {
	purchase.Products = append(purchase.Products, *p)
}

func (purchase ShoppingCartImpl) GetTotal() float32 {
	var total float32
	for _, product := range purchase.Products {
		total = total + product.Price
	}
	return total
}

func (purchase ShoppingCartImpl) GetTotalWithDiscount(strategy discount.DiscountStrategy) float32 {
	total := purchase.GetTotal()
	return strategy.GetDiscount(total)
}
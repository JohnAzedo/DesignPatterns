package main

import (
	"DesignPatterns/strategy"
	"DesignPatterns/strategy/discount"
	"fmt"
)

func main(){

	disc := discount.NewDiscountBlackFriday()
	shoppingCart := strategy.NewShoppingCart()
	shoppingCart.AddProduct(&strategy.Product{Name: "TV", Price: 100.0})
	shoppingCart.AddProduct(&strategy.Product{Name: "TV", Price: 100.0})
	fmt.Println(shoppingCart.GetTotalWithDiscount(disc))
}
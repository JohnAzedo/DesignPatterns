package factory

import "fmt"

type  IHeroShip interface {
	IShip
	ShowSloganHeroShip()
}

type HeroShip struct {
	Ship
}

func NewHeroShip(title string) IHeroShip {
	return &HeroShip{
		Ship: Ship{title, 20, 100},
	}
}

func (ship HeroShip) ShowSloganHeroShip() {
	fmt.Println(ship.Title, "says: To the infinity and beyond.")
}
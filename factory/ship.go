package factory

import "fmt"

type IShip interface {
	Shot(target IShip)
	OnScreen()
	Hit(damage int)
	GetTitle() string
}

type Ship struct {
	Title string;
	Damage int;
	Life int;
}

func (ship Ship) GetTitle() string {
	return ship.Title
}

func (ship Ship) Shot(target IShip) {
	fmt.Println(ship.Title, "shot", target.GetTitle(), "!")
	target.Hit(ship.Damage)
}

func (ship Ship) OnScreen() {
	fmt.Println(ship.Title, "has", ship.Life, "life left")
}

func (ship *Ship) Hit(damage int) {
	ship.Life -= damage
	fmt.Println(ship.Title, "was hit!")
	ship.OnScreen()

	if ship.Life <= 0 {
		fmt.Println(ship.Title, "sank!")
	}
}
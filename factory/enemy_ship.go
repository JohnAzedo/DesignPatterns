package factory

import "fmt"

type IEnemyShip interface {
	IShip
	ShowBossName()
}

type EnemyShip struct {
	Ship
}

func NewEnemyShip(title string) IEnemyShip {
	return &EnemyShip{
		Ship{title, 15, 100},
	}
}

func (ship EnemyShip) ShowBossName() {
	fmt.Println(ship.Title, " belongs to Vader")
}
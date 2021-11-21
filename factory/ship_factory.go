package factory

import (
	"log"
)

type ShipFactory struct {}

func (factory ShipFactory) Make(shipType string, shipTitle string) IShip {
	if shipType == "Hero" {
		return NewHeroShip(shipTitle)
	} else if shipType == "Enemy" {
		return NewEnemyShip(shipTitle)
	} else {
		log.Fatal("ShipType %v is not allowed!", shipType)
		return nil
	}
}
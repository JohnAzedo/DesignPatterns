package main

import (
	. "DesignPatterns/factory"
)

func main(){
	shipFactory := ShipFactory{}
	heroShip := shipFactory.Make("Hero", "(Hero) ShipOne").(IHeroShip)
	enemyShip := shipFactory.Make("Enemy", "(Enemy) ShipTwo")

	heroShip.Shot(enemyShip)
	heroShip.ShowSloganHeroShip()
}
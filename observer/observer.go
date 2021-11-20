package observer

import "fmt"

type IObserver interface {
	update()
}

type Observer struct {
	ID int
}

func (observer Observer) update(){
	fmt.Println("Observer ", observer.ID, " running!")
}

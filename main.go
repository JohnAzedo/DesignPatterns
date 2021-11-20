package main

import (
	. "DesignPatterns/observer"
)

func main(){
	subject := Subject{}

	subject.Subscribe(Observer{1})
	subject.Subscribe(Observer{2})
	subject.Subscribe(Observer{3})

	subject.SetState(2)
}
package observer

import "errors"

type ISubject interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	SetState(arg interface{})
	NotifyAll()
}

type Subject struct {
	observers []Observer
	state     interface{}
}

func (subject *Subject) SetState(arg interface{}){
	subject.state = arg
	subject.NotifyAll()
}

func (subject Subject) GetObserversAmount() int{
	return len(subject.observers)
}

func (subject* Subject) Subscribe(observer Observer) {
	subject.observers = append(subject.observers, observer)
}

func (subject* Subject) Unsubscribe(observer Observer){
	position, err := subject.indexOf(observer)
	if err != nil {
		subject.observers = append(subject.observers[:position], subject.observers[position+1:]...)
	}
}

func (subject Subject) NotifyAll() {
	for _, obs := range subject.observers {
		obs.update()
	}
}

func (subject Subject) indexOf(observer Observer) (int, error){
	for i, obs := range subject.observers {
		if obs == observer {
			return i, nil
		}
	}
	return -1, errors.New("Observer is not on array!")
}




package observer

import (
	"reflect"
	"testing"
)

func TestObserver(t *testing.T){
	subject := Subject{ state: 1 }

	subject.Subscribe(Observer{1})
	subject.Subscribe(Observer{2})
	subject.Subscribe(Observer{3})

	expected := 3
	result := subject.GetObserversAmount()

	if !reflect.DeepEqual(result, expected){
		t.Fatalf("Error: %v (result) is not the same as %v (expected)", result, expected)
	}
}
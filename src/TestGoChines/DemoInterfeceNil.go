package main

type I interface {
	M()
}

type T struct {
	S string
}

func (t T)M(){
	if t == nil {

	}
}
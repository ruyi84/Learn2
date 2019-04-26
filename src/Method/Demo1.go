package main

import "fmt"

type address struct{
	city string
	state string
}

func (a address) funllAddress(){
	fmt.Printf("Full address:%s,%s",a.city,a.state)
}

type person struct{
	firstName string
	lastName string
	address
}

func main(){
	p:= person{
		firstName:"Elon",
		lastName:"Must",
		address:address{
			city:";Los Angeles",
			state:"California",
		},
	}

	p.funllAddress()
}
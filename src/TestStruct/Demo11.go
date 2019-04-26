package main

import "fmt"

var  ss struct{
	id int
	name string
	a func()

}

type Address struct {
	city, state string
}
type Person struct {
	name string
	age  int

	Address
}

func main(){
	ss.a = func() {
		fmt.Println("wwww")
	}
	man := Person{"beijing",11,"asd"}
	fmt.Println(man)

}

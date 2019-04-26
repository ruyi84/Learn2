package main

import "fmt"

var  SuperHero struct{
	age int
	name string
	SetAge func(int)
	SetName func(string)
	tostring func()
}



func main(){
	SuperHero.SetAge = func(age int) {
		SuperHero.age = age
	}
	SuperHero.SetName = func(name string) {
		SuperHero.name = name
	}
	SuperHero.tostring = func(){
		fmt.Println("Age:",SuperHero.age,"name",SuperHero.name)
	}
	SuperHero.SetName("Batman")
	SuperHero.SetAge(11)
	SuperHero.tostring()
}

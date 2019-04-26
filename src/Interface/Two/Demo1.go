package main

import "fmt"

type Describer interface {
	Describer()
}

type Person struct {
	name string
	age int
}

func (p Person) Describer(){
	fmt.Printf("%s is %d years old\n",p.name,p.age)
}

type Address struct{
	state string
	country string
}

func (a *Address) Describer(){
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main(){
	var d1 Describer
	p1 := Person{"Sam",25}
	d1 = p1
	d1.Describer()
	p2 := Person{"james",32}
	d1 = &p2
	d1.Describer()

	var d2 Describer
	a := Address{"Washington","USA"}
	d2 = &a
	d2.Describer()
}
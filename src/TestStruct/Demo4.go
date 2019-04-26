package main

import "fmt"

var name1 struct{
	Name string
	age func() int
}

func main(){



	var n = name1
	var n1 = name1
	fmt.Println(n)
	fmt.Println(n1)
}
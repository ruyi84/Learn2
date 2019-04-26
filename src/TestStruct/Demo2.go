package main

import "fmt"


func main(){

	//仅仅是定义了一个结构体变量，而没有定义任何结构体类型
	emp3 := struct {
		firstName,lastName string
		age,salary int
	}{
	firstName: "Andreah",
		lastName:"Nikola",
			age: 31,
			salary: 5000,
	}




	fmt.Println(emp3)
}
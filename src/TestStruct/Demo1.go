package main

import "fmt"

type Employee struct{
	firstName,lastName string
	age,salary int
}

func main(){
	emp1 :=Employee{
		firstName: "sam",
		age: 25,
		salary : 500,
		lastName: "Anderson",
	}

	emp2 := Employee{"thomas","paul",29,800}



	fmt.Println(emp1)
	fmt.Println(emp2)
}

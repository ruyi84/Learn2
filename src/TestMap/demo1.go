package main

import "fmt"

func main(){
	personSalary := map[string]int{
		"steve" : 12000,
		"jamie" : 15000,
	}
	personSalary["mike"] = 9000
	newEmp :="joe"
	value,ok := personSalary["mike"]
	if ok == true{
		fmt.Println(newEmp,"is",value)
	}else{
	fmt.Println(newEmp,"not found",value)
}
}
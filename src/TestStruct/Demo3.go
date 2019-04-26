package main

import "fmt"

var user struct{
	Name string
	age int
}



func main(){
	user.Name = "name"
	user.age = 18

	fmt.Println(user)



}
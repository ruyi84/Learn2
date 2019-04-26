package main

import "fmt"

var a struct{
	int
}

var b struct{
	int
}

func main(){
	fmt.Println(a == b)
}

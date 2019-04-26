package main

import "fmt"

func main(){
	var names []string
	if names == nil{
		fmt.Printf("slice is nil going to append")
		names = append(names,"johe","sebastian","vinay")
		fmt.Println("names contents:",names)
	}

}

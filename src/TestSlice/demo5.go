package main

import "fmt"

func subtactOne(numbers []int){
	for i,_ := range numbers{
		numbers[i] -=2
	}
}

func main(){
	nos := []int{8,7,6}
	fmt.Println(nos)
	subtactOne(nos)
	fmt.Println(nos)
}
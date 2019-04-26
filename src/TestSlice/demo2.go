package main

import "fmt"

func main(){
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n",len(fruitslice),cap(fruitslice))
	fruitslice = fruitslice[:cap(fruitslice)]
	fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
}
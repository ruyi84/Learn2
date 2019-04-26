package main

import "fmt"

type num struct {
	name string
	ai func()
}


func main(){
	s := []int{1,2}
	slice_attribute(s)
	s = append(s,4,5,6)
	slice_attribute(s)
	s = append(s,2,3,3)
	slice_attribute(s)
}

func slice_attribute(s []int){
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}
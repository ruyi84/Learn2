package main

import "fmt"

func main(){
	var s []int
	s =[]int{1,2,3,4}
	slice_attribute(s)
	s = append(s,0)
	slice_attribute(s)
	s = append(s,2,3,4,3,4,4,1,3)
	slice_attribute(s)
}

func slice_attribute(s []int){
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}
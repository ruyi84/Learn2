package main

import "fmt"

func mkslice() []int{
	s := make([]int,1,2)
	s =append(s,100)

	return s
}

func mkmap() map[string]int{
	m := make(map[string]int)
	m["a"] =1
	return m
}

func main(){
	//m :=mkmap()
	//println(m["a"])

	s :=mkslice()
	println(s[0])
	fmt.Println(s)
	fmt.Println(s[0])


}

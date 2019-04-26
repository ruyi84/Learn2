package main

import "fmt"

func main(){
	arr := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(arr)
	var sum []int = arr[1:3]
	fmt.Println(sum)

	fmt.Printf("%d\n",&arr)
	fmt.Printf("%d\n",&sum)

	num := []int{1,2,3}
	var s []int =[]int{1,1,1,1,1}
	fmt.Println("s len and cap :",len(s),cap(s))
	fmt.Println("num :",num)
	fmt.Println("s:",s)
	s = num[:2]
	fmt.Println("s:",s)
	fmt.Println("s len and cap :",len(s),cap(s))
}
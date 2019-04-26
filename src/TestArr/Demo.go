package main

import "fmt"

func main(){
	var a [3] int
	a[0] =12
	a[1] =112
	a[2] =121
	b :=[...]int{1,2,31}
	fmt.Println(len(b))
}
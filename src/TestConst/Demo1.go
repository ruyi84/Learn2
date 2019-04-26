package main

import "fmt"


func main(){
	//常亮是无类型的，但是无类型的常量有一个关联的默认类型，并且仅当一行代码需要时才提供

	const a = "aaa"

	const a2 string = "ma"
	b := a
	b2 := a2
	var b3 string = a
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",b2)
	fmt.Println(b3)
}

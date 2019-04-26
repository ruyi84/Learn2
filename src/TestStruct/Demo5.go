package main

import (
	"fmt"
	"unsafe"
)

func main(){
	array1 := [4]int{1,2,3} //定义array2这个数组，开辟了一款内存。
	fmt.Printf("array1的元素是：%d\n",array1)
	fmt.Printf("array1数组所占内存是:%d bytes\n",unsafe.Sizeof(array1)) //一个数组占有8个字节，容量为4的数组其内存是就是32字节
	var  array2 [4]int  //定义一个
	array2 = array1
	fmt.Printf("array1的地址是：%d\narray2的地址是：%d\n",&array1[0],&array2[0])
	fmt.Println(array1,array2)
	array1[0] = 3
	fmt.Println(array1,array2)
}
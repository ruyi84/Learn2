package main

import "fmt"

//测试关于指针方法与指针函数
//带指针参数的函数，参数只能接受一个指针。而以指针为接受者的方法被调用时，接受者既能为值又能为指针。

type Man struct {
	name string
	age int
}

type Super interface {
	S()
}

func (m Man)S(){
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaa")
}

func toString(m *Man){
	println("Nmae :" , m.name,"\nAge :", m.age)
}

func (m *Man)toString(){
	println("Nmae :" , m.name,"\nAge :", m.age)
}

func main(){
	m := Man{"Bat man" , 32}
//	mm := &m

//	toString(&m)
//	m.toString()
//	mm.toString()

	fmt.Printf("%T\n",m)
	fmt.Println(m)
}
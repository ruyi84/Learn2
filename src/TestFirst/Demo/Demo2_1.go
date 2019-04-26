package Demo

import "fmt"

func simple1() func(a,b int) int{
	f := func(a,b int) int{
		return a+b
	}
	return f
}

func main(){
	s := simple1()
	fmt.Println(s(50,8))
}

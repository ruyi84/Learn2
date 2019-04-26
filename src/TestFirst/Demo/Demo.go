package Demo

import "fmt"

func appendStr() func(string) string{
	t := "Hello"
	fmt.Println(t)
	c := func(b string) string{
		t = t + " " + b
		return t
	}

	return c
}

func main(){
	a := appendStr()

	fmt.Println(a("World"))
	fmt.Println(a("World"))
}

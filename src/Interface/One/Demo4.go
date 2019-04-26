package One

import "fmt"

func assert(i interface{}){
	v,ok := i.(int)
	fmt.Println(v,ok)

}

func main(){
	var s interface{}= 51
	assert(s)
	var i interface{} ="sdasd"
	assert(i)

}

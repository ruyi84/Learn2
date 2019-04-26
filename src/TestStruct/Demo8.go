package main

import "fmt"

var cmd string

func main(){
	for {
		fmt.Println("input>>")
		fmt.Scan(&cmd)
		if cmd == "stop"{
			break
		}
		fmt.Println(cmd)
	}
}
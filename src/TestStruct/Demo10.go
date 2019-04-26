package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Student struct {
	ID int
	Name string
}

func main(){
	s := Student{
		ID :1,
		Name:"yijing",
	}
	buf,err := json.Marshal(s)
	if err != nil{
		log.Fatal(err)
	}
	f,err :=os.OpenFile("F:\\123.txt",os.O_APPEND|os.O_CREATE|os.O_RDWR,0644)
	if err !=nil{
		log.Fatal(err)
	}
	f.WriteString(string(buf))
	f.WriteString("\n")
	f.Close()
	fmt.Println("Win")
}
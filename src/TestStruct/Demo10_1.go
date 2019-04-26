package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Student1 struct{
	ID int
	Name string
}

func main(){
	buf,err :=ioutil.ReadFile("F:\\123.txt")
	if err!=nil{
		fmt.Println("Low")
		return
	}
	var str Student1
	err1 :=json.Unmarshal(buf,&str)
	if err1 !=nil{
		log.Fatal("111111",err1)
	}
	fmt.Println(str)
}
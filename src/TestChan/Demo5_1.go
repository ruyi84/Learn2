package main

import (
	"fmt"
	"time"
)

func write(ch chan int){
	for i:=0;i <5;i++{
		ch <-1
		fmt.Println("successfully wrote",i,"to ch")
	}
	close(ch)
}
func main(){
	ch := make(chan int,2)
	go write(ch)
	time.Sleep(2*time.Second)
	for v := range ch{
		fmt.Println("rede value",v,"from ch")
		time.Sleep(2*time.Second)
	}
}
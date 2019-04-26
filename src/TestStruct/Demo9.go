package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var(
	s string
	n int
	line string
)

func main(){
	f := bufio.NewReader(os.Stdin)
	for{
		fmt.Print("请输入一些字符串>")
		line,_ = f.ReadString('\n')
		if len(line) == 1{
			continue
		}
		line = strings.Replace(line,"\n"," ",-1)
		fmt.Printf("您输入的是:%s\n",line)
		if s == "stop"{
			break
		}
		fmt.Printf("您输入的第一个参数是：·\033[31;1m%v\033[0m·,输入的第二个参数是··\033[31;1m%v\033[0m·.\n",s,n)
	}
}
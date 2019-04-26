package main

import (

	"fmt"
	"strconv"
)

type IPAddr [4] byte

func (i IPAddr)String() string{
	l := len(i)-1
	var s string
	for index,v := range i{
		s += strconv.Itoa(int(v))
		if index != l {
			s += "."
		}
	}
	return s
}

func main(){
	hosts := map[string]IPAddr{
		"loopbakc":{127,0,0,1},
		"googleDNS":{8,8,8,8},
	}

	for name,ip := range hosts{
		fmt.Printf("%v:%v\n",name,ip)
	}

}
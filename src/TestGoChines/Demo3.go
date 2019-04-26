package main

import (
	"sync"
)

func main(){
	wg := sync.WaitGroup{}

	si := []int{1,2,3,4,5,6,7,8,9}

	for i := range si {
		wg.Add(1)

		go func() {
			println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

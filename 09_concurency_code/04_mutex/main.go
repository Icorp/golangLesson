package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu sync.Mutex
	counter = 0
	)

func main() {
	for i:=0;i<10;i++{
		go incr()
	}
	time.Sleep(time.Millisecond*10)
}
func incr()  {
	mu.Lock()
	defer mu.Unlock()
	counter++
	fmt.Println(counter)
}

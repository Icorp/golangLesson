package main

import (
	"fmt"
	"sync"
)
 var wg sync.WaitGroup
func main()  {
	wg.Add(2)
	go worker()
	go worker2()
	wg.Wait()
}

func worker()  {
	for i:=0;i<10;i++{
		fmt.Println("Worker#1:",i)
	}
	wg.Done()
}
func worker2()  {
	for i:=0;i<10;i++{
		fmt.Println("Worker#2:",i)
	}
	wg.Done()
}



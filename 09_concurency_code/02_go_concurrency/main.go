package main

import (
	"fmt"
)

func main() {
	go worker()
	go worker2()
	fmt.Scanln()
}

func worker() {
	for i := 0; i < 10; i++ {
		fmt.Println("Worker#1:", i)
	}
}
func worker2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Worker#2:", i)
	}
}

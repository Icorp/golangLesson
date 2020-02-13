package main

import (
	"fmt"
)

func main() {
	even:=make(chan int)
	odd:=make(chan int)
	quit:=make(chan bool)
	go send(even,odd,quit)
	receiver(even,odd,quit)
	fmt.Println("Finish")
}
func receiver(even,odd <-chan int,quit <-chan bool)  {
	for  {
		select {
		case v:=<-even:
			fmt.Println("Четные числа:",v)
		case v:=<-odd:
			fmt.Println("Нечетные числа",v)
		case v:=<-quit:
			fmt.Println("Кейс для выхода",v)
			return
		}
	}
}

func send(even, odd chan<- int, quit chan bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <-true
}

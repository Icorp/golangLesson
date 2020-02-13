package main

import "fmt"

func main() {
	c:=make(chan int)
	go send(c)
	receiver(c)
}
func send(c chan<- int)  {
	c <-40
}
func receiver(c <-chan int)  {
	fmt.Println(<-c)
}
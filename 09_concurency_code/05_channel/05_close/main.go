package main

import "fmt"

func main()  {
	c:=make(chan int)
	close(c)
	go send(c)
	receiver(c)
	fmt.Println("Finish")
}
func send(c chan<-int)  {
	c<-42
}
func receiver(c <-chan int)  {
	fmt.Println("Value = ",<-c)
}



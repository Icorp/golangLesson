package main

import "fmt"

func main() {
	x:=10
	worker(&x)
	fmt.Println(x)
}
func worker(x *int)  {
	*x=100
	fmt.Println(x)
}

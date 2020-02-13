package main

import "fmt"

func main() {
	a,b:=sayHello("James","Bond")
	fmt.Println(a)
	fmt.Println(b)
}
func sayHello(fn string,ln string)(string,bool)  {
	a:= fmt.Sprint(fn," ",ln, " Hello")
	b:= true
	return a,b
}
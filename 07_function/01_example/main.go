package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1:=woo("James")
	fmt.Println(s1)
}

func foo()  {
	fmt.Println("Hello from foo")
}
func bar(s string)  {
	fmt.Println("Hello",s)
}
func woo(st string)string  {
	return fmt.Sprint("Hello from woo",st)
}
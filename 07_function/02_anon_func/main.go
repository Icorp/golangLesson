package main

import "fmt"

func main() {
	foo()
	func(){
		fmt.Println("Anonymous func run")
	}()
	func(x int){
		fmt.Println("Your number is",x)
	}(42)
}
func foo()  {
	fmt.Println("foo run")
}

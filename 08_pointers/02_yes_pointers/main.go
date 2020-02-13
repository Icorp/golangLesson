package main

import "fmt"

func main()  {
	x:=0
	fmt.Println("x до ", &x)
	fmt.Println("x до ", x)
	foo(&x)
	fmt.Println(x)
}
func foo(y *int)  {
	fmt.Println("y до ", y)
	fmt.Println("y до ", *y)
	*y = 43
	fmt.Println("y после ", y)
	fmt.Println("y после", *y)
}

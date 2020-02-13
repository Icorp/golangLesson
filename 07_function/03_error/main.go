package main

import "fmt"

func main()  {
	a,err:=fmt.Println("Hello")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(a)
	fmt.Println(err)
}
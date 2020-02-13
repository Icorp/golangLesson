package main

import "fmt"

func main() {
	m:= map[string]int{
		"Andrey":25,
		"Vasya":27,
	}
	fmt.Println(m)
	m["key"] = 33
	fmt.Println(m)
	delete(m,"key")
	fmt.Println(m)
	fmt.Println(m["key"])
}

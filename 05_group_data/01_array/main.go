package main

import "fmt"

func main() {
	var array [5] int
	fmt.Println(array)
	array[3]=15
	fmt.Println(array)
	fmt.Println(len(array))
}

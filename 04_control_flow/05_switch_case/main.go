package main

import "fmt"

func main() {
	x := 40
	switch x {
	case 41:
		fmt.Println("41")
	case 42:
		fmt.Println("42")
	case 43:
		fmt.Println("43")
	case 44:
		fmt.Println("44")
	default:
		fmt.Println("I dont know")
	}

	fmt.Println("Finish")
}

package main

import "fmt"

func main() {
	names := []string{"nvidia","amd"}
	fmt.Println(names)
	names2 := []string{"intel","tesla"}
	fmt.Println(names2)
	xp:=[][]string{names,names2,}
	fmt.Println(xp)
}

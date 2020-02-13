package main

import (
	"fmt"
	"os"
)

func main() {
	file,err:=os.Create("10_working_with_file/01_create_file/file.txt")
	if err != nil{
		fmt.Println("Something is wrong")
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(file.Name())
}

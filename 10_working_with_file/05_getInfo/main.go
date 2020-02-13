package main

import (
	"fmt"
	"os"
)

func main() {
	fileStat, err := os.Stat("10_working_with_file/05_getInfo/file.txt")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println("File Name:",fileStat.Name())
	fmt.Println("File Size:",fileStat.Size())
	fmt.Println("File permissions:",fileStat.Mode())
	fmt.Println("Last modified:",fileStat.ModTime())
}

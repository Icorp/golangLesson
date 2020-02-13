package main

import (
	"fmt"
	"os"
)

func main() {
	data:=[]byte("Hello")
	file,err:=os.Create("10_working_with_file/03_write_file/file.txt")
	if err != nil{
		fmt.Println("Something is wrong")
		os.Exit(1)
	}
	defer file.Close()
	_, _ = file.Write(data)
	fmt.Println("Done.")
}

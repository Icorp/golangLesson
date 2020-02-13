package main

import (
	"fmt"
	"os"
)

func main() {
	err:=os.Remove("10_working_with_file/03_write_file/file.txt")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("File deleted")
}

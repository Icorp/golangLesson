package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data,err:=ioutil.ReadFile("10_working_with_file/02_read_file/text.txt")
	if err!=nil{
		fmt.Println("File reading error",err)
		return
	}
	fmt.Println("Content of file:",string(data))
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.ru")
	if err !=nil{
		fmt.Println(err)
	}
	_, _ = io.Copy(os.Stdout, resp.Body)
}

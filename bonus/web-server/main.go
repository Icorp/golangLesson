package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/",handler)
	fmt.Println("Server starting on 8000")
	_ = http.ListenAndServe(":8000", nil)
}
func handler(w http.ResponseWriter,r *http.Request)  {
	_, _ = w.Write(sendRequest())
}

func sendRequest()[]byte  {
	resp,err:=http.Get("http://yandex.ru")
	if err !=nil{
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	f,err:=ioutil.ReadAll(resp.Body)
	if err !=nil{
		fmt.Println(err)
	}
	return f
}
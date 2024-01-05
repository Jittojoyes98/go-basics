package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("Post request here")
	PerformPostRequest()
}

func PerformPostRequest(){
	const myurl string ="https://jsonplaceholder.typicode.com/posts"

	requestBody:=strings.NewReader(`
	{
		"title": "foo",
		"body": "bar",
		"userId": 1
	}`)

	response,err:=http.Post(myurl,"application/json",requestBody)

	if err!=nil{
		panic(err)
	}
	content,err:=io.ReadAll(response.Body)
	fmt.Println(string(content))

}
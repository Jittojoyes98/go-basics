package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("Get request")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl="https://lco.dev/"
	response,err:=http.Get(myurl)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code for the request",response.StatusCode)
	fmt.Println("Content length",response.ContentLength)

	var responseString strings.Builder
	content,err:=ioutil.ReadAll(response.Body)
	byteCount,_:=responseString.Write(content)
	fmt.Println(byteCount,responseString.String())
	// fmt.Println("the data ",string(content))

}
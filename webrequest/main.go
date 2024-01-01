package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url="https://lco.dev"

func main()  {
	fmt.Println("LCO request starts here..")
	response,err:=http.Get(url)
	if err!=nil{
		panic(err)
	} 
	defer response.Body.Close() // we need to close them thats our responsibility
	fmt.Printf("The response type here is ... %T \n",response)

	databytes,err:=ioutil.ReadAll(response.Body)

	if err!=nil{
		panic(err)
	} 
	fmt.Println("The response data here is nice and easy ", string(databytes))
}

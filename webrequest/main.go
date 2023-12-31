package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const requestUrl="https://lco.dev"

func main()  {
	fmt.Println("LCO request starts here..")
	response,err:=http.Get(requestUrl)
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


	// usage of url package here.
	
	u,err:=url.Parse("http://bing.com/search?q=dotnet")
	if err!=nil{
		panic(err)
	}
	
	fmt.Println("Curent scheme is ", u.Scheme)
	u.Host="google.com"
	q:=u.Query()
	q.Set("q","golang") // changing the query here.
	fmt.Println("The query is ",q)
	fmt.Println("the updated url",u)
}

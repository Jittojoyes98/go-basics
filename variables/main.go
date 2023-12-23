package main

import "fmt"

// global is caps
const LoginToken string ="token"

func main()  {
	var username string = "hello"
	// fmt.Println("HEEY") 
	const isLoggedIn bool = true
	fmt.Print(isLoggedIn,"\n")
	fmt.Printf("The current user is %T \n",username)

	var password="abcdef"
	fmt.Print(password,"\n")

	// modulus operator
	newUser:= "hwwy"
	fmt.Print(newUser,"\n")
	fmt.Println(LoginToken)
}

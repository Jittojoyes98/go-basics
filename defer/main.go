package main

import "fmt"

func main()  {
	defer fmt.Println("WORLD")
	defer fmt.Println("HELLO")
	fmt.Println("Heyy,")
	orderPrint()
}

// remember it works as a stack 
// WORLD,HELLO, 01234
func orderPrint()  {
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}
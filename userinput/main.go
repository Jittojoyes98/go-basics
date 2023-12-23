package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Hello Welcome user "
	fmt.Println(welcome)
	
	reader:=bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating")

	rating,_:=reader.ReadString('\n')

	fmt.Println("The rating for the pizza is ",rating)

}
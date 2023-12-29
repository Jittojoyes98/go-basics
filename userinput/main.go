package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main()  {
	welcome := "Hello Welcome user "
	fmt.Println(welcome)
	
	reader:=bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating")

	rating,_:=reader.ReadString('\n')

	fmt.Println("The rating for the pizza is ",rating)

	presentTime:=time.Now()

	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	// always in this format is crazy.

}
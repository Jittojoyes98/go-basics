package main

import "fmt"

func main()  {
	fmt.Println("Heey")
	var ans int=adder(1,3)
	println(ans) 
	 result,message := calculate(1,2,3,5)
	println(result,message)
}

// specify the types and return type
func adder(a int,b int) int {
	return a+b
}

func calculate(sumarray ...int) (int,string)  {
	total:=0
	for _,val := range(sumarray){
		total+=val
	}
	return total,"HeY THERE"
}

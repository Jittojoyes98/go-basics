package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println("File operations in Go")
	file,err:=os.Create("./createdFile.txt")
	handleError(err)
	const content string="Hey there new file is here."
	length,err:=file.WriteString(content)
	handleError(err)
	fmt.Print("The file length is \n",length)
}

func handleError(errorString error)  {
	if errorString!=nil{
		panic(errorString)
	}
}
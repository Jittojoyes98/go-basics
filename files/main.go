package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("File operations in Go")
	file,err:=os.Create("./createdFile.txt")
	handleError(err)
	const content string="Hey there new file is here."
	length,err:=file.WriteString(content)
	// here you can write like :
	// os.WriteString(file,content)
	handleError(err)
	fmt.Print("The file length is \n",length)

	// Read files using ioutil 
	readFile("./createdFile.txt")
}

func readFile(filePath string)  {
	databyte,err:=ioutil.ReadFile(filePath)
	handleError(err)
	print("\n")
	fmt.Println("Text data in the file:",string(databyte))

}

func handleError(errorString error)  {
	if errorString!=nil{
		panic(errorString)
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type User struct  {
	name string
	RollNo int
}

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

	var courses= []string{"Math","Chem","Phy","Comp"} // not having a fixed size
	// courses[0]="Math"
	fmt.Println(courses)
	var index int = 2
	fmt.Println(index)
	// to remove a course on index 2.
	courses = append(courses[:index], courses[index+1:]...) 
	fmt.Println(courses)

	mapping:=make(map[string]string)

	mapping["JS"]="Javascript"
	mapping["RB"]="Rubi"
	mapping["PY"]="Python"

	fmt.Println("Hey ther the variable",mapping)
	fmt.Println("Hey the key value is",mapping["JS"])
	ans:=User{"Jitto",18}
	fmt.Println(ans)

	fmt.Printf("New user is %+v \n",ans)

	// for d:=0; d<len(courses); d++{
	// 	fmt.Println(courses[d])
	// }
	for _,val := range(courses){
		fmt.Println(val)
	}

}


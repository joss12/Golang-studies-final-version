package basics

import "fmt"


func main (){

	//Use cases of range
	message := "Hello World"
	for  i, v := range message{
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
		
	}

}
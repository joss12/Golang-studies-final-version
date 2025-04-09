package basics

import "fmt"


func main (){

	//Panic (interface{})

	//example of valid input
	process(10)
	defer fmt.Println("Deferred 1")

	//Example of an invalid input
	process(-3)
}

func process(input int){

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0{
		fmt.Println("Before Panic")
		panic("input must be a non--negative number")
		// fmt.Println("After Panic");
		// defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}
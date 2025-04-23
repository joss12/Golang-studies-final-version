package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Command:", os.Args[0])
	for i, arg := range os.Args {
		fmt.Println("Argument1:", i, ":", arg)
	}

	//Define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "Eddy", "Name of the user")
	flag.IntVar(&age, "age", 22, "The age of the user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	//Parsing the command line arg
	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)

}

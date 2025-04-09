package basics

import (
	"fmt"
	"os"
)


func main (){
	fmt.Println("Starting the main function")

	//Exit with status code of 1
	os.Exit(1)

	//Thee will never be executed
	fmt.Println("end of the main function")

}
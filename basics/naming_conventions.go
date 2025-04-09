package basics

import "fmt"

type employee struct {
	firsName string
	lastName string
	Age      int
}


type employeeApple struct {
	firsName string
	lastName string
	Age      int
}
func main() {

	const MAXETRIES = 5
	var employeeID = 1001
	fmt.Println("EmployeeID:", employeeID)
}

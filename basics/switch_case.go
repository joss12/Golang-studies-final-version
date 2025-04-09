package basics

import "fmt"

var pl = fmt.Println

func main() {

	//switch statement
	// switch expression{
	// case value1:
	// 	code to be executed if expression equals value1
	// case value2:
	// 	code to be executed if expression equals value2
	// case value3:
	// 	code to be executed if expression equals value3
	// default:
	// 	code to be executed if expression does not match any value
	// }

	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's a banana.")
	default:
		fmt.Println("Unknown Fruits!")
	}

	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	pl("It's a weekday.")
	// case "Sunday":
	// 	pl("It's a weekend.")
	// default:
	// 	pl("Invalid day.")
	// }

	// number := 15

	// switch {
	// case number < 10:
	// 	pl("Number is less than 10")
	// case number >= 10 && number < 20:
	// 	pl("Number is between 10 and 19")
	// default:
	// 	pl("Number is 20 or more")
	// }

	// num := 2
	// switch {
	// case num > 1:
	// 	pl("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	pl("Number is 2")
	// default:
	// 	pl("Not Two")
	// }
	
	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x any){
	switch x.(type){
	case int:
		pl("Integer");
	case int32:
		pl("It's an integer 32")
	case float64:
		pl("It's a float")
	case string:
		pl("It's string")
	default:
		pl("Unknown type")
	}
}

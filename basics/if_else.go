package basics

import "fmt"

func main() {

	// condition

	// age := 25

	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// }

	//if condition{
	//block of code
	// }else if{
	//block of code
	// }else{
	// block of code
	// }

	// temperaturte := 25
	// if temperaturte >= 30{
	// 	fmt.Println("It's hot outside")
	// }else{
	// 	fmt.Println("It's cold outside")
	// }

	score := 85

	// if score >= 90{
	// 	fmt.Println("Grade A")
	// }else if score >= 80{
	// 	fmt.Println("Grade B")
	// }else if score >= 70{
	// 	fmt.Println("GRADE C")
	// }else{
	// 	fmt.Println("GRADE D")
	// }

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("GRADE C")
	} else {
		fmt.Println("GRADE D")
	}

	// this line will be executed after one the condition is met

	/*
		if condition{
		code block1
		if condition2{
		code block 2
		}
		}
	*/

	// num := 15
	// if num % 2 == 0{
	// 	if num % 3 == 0{
	// 		fmt.Println("Number is divisible by both 2 and 3.")
	// 	}else{
	// 		fmt.Println("Number is divisible by 2 but not 3")
	// 	}
	// }else{
	// 	fmt.Println("Number is not divisible by 2.")
	// }

	// || OR
	// && AND

	if 11%2 == 0 || 6%2 == 0 {
		fmt.Println("Either 10 or 5 are even.")
	}

	if 10%2 == 0 || 6%2 == 0 {
		fmt.Println("Both 10 or 6 are even.")
	}

}

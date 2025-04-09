package basics

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Reminder:", result)

	const p float64 = 22.0 / 7
	fmt.Println(p)


	//Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // max value that int64 can can hold
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	//Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // max value for unit64 ty
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	//underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}

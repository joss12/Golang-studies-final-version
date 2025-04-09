package basics

import "fmt"

func main() {
	// var arrayName [size]elementType
	// var numbers [5]int
	// fmt.Println(numbers)

	// numbers[4] = 20
	// fmt.Println(numbers)

	// numbers[0] = 9
	// fmt.Println(numbers)

	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	// fmt.Println("Fruits Array:", fruits)

	// fmt.Println("Third element:", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// copieArray := originalArray

	// copieArray[0] = 10
	// fmt.Println("Original Array:", originalArray)
	// fmt.Println("Copied Array:", copieArray)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Element at index", i, ":", numbers[i])
	// }

	// //underscore is blank identifier, used to store unused values
	// for _, v := range numbers {
	// 	fmt.Printf(" Value: %d\n", v)
	// }

	// fmt.Println("The length of the numbers array is", len(numbers))

	// //Comparing Arrays
	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{1, 2, 3}

	// a, _ := someFunction()
	// fmt.Println(a)
	// //  fmt.Println(b)

	// b := 2
	// _ = b

	// fmt.Println("Array1 is equal to Array2", array1 == array2)
	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copieArray *[3]int

	copieArray = &originalArray
	copieArray[0] = 100
	fmt.Println("Original Array:", originalArray)
	// fmt.Println("Copied Array:", *copieArray)

}

func someFunction() (int, int) {
	return 1, 2
}

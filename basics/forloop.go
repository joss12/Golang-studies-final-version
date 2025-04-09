package basics

import "fmt"

func main() {

	// Simple iteration over range
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	//iteration over collection
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("ood Number:", i)
	// 	if i == 5 {
	// 		break //break out of the loop
	// 	}
	// }


	//ASTERIK LAYOUT
	// rows := 5

	// // outer loop
	// for i := 1; i<=rows;i++{
	// 	// inner loop for space before starts
	// 	for j:=1;j<=rows-i;j++{
	// 		fmt.Print(" ")
	// 	}
	// 	for k:=1;k<=2*i-1;k++{
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // Move to the next line
	// }

	for i := range 10{
		i++
		fmt.Println(i)
	}
	fmt.Println("We have lift off")
}

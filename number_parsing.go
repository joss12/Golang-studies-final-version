package main

import (
	"fmt"
	"strconv"
)


func main (){
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil{
		fmt.Println("Enter parsing the value:", err)
	}
	fmt.Println("Parsed Integer:", num)
	fmt.Println("Parsed Integer:", num+1)
	
	numiStr, err :=  strconv.ParseInt(numStr, 10, 32)
	if err != nil{
		fmt.Println("Enter parsing the value:", err)
	}
	fmt.Println("Parsed Integer", numiStr)

	floatstr := "3.14"
	floatval, err:= strconv.ParseFloat(floatstr, 64)
	if err != nil{
		fmt.Println("Enter parsing value:", err)
	}
	fmt.Printf("Parsed float: %.2f\n", floatval)

	binaryStr := "1010" // 0 + 2 + 0 + 8
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil{
		fmt.Println("Error parsing binary value:", err)
		return
	}
	fmt.Println("Parsed binary to decimal:", decimal)

	hexStr := "FF" // 0 + 2 + 0 + 8
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil{
		fmt.Println("Error parsing binary value:", err)
		return
	}
	fmt.Println("Parsed hex to decimal:", hex)


	invalidNum := "456abc" // 0 + 2 + 0 + 8
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil{
		fmt.Println("Error parsing value:", err)
		return
	}
	fmt.Println("Parsed hex to decimal:", invalidParse)
}
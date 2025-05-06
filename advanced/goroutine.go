package main

import (
	"fmt"
	"time"
)

/*
Goroutines are just functions that leave the main thread
run in the background and come back to join the main thread
once the functions are finished/ready to return any value
Goroutine do mot stop the program flow and non blocking
Goroutine are mostly anonymous functions
*/

func main() {
	//Error variable
	var err error

	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After sayHello function.")

	go func() {
		err = doWork()
	}()

	// err = go doWork() // this is not acceptable
	go printNumbers()
	go printLetters()
	if err != nil {
		fmt.Println("Error:", err)
	}

	//basic for a goroutine to finish
	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully.")
	}

}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 6; i++ {
		fmt.Println("Number:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	//Similate work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occurred in doWork.")
}

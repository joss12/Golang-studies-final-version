package advanced

import (
	"fmt"
	"time"
)

func unbuffered() {
	// buffer channel means a channel with storage  a buffer is a storage
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		// fmt.Println(<-ch)
		fmt.Println("3 second Goroutine finished")
		// time.Sleep(3 * time.Second)
	}()

	ch <- 1

	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(3*time.Second)
	// 	fmt.Println("3 second Goroutine finished")

	// }()

	// receiver = <-ch
	// fmt.Println(receiver)
	fmt.Println("End of program.")
}

package advanced

import (
	"fmt"
	"time"
)

// func main() {
// 	// done := make(chan struct{})
// 	// done := make(chan bool)
// 	done := make(chan int)

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		// done <- struct{}{}
// 		// done <- false
// 		done <- 0
// 	}()

// 	<-done
// 	fmt.Println("Finished.")
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		fmt.Println("Sending...")
// 		ch <- 9 // Blocking until the value is received
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Sent value")
// 	}()

// 	value := <-ch // Blocking until a value is sent
// 	fmt.Println(value)
// }

// =========== SYNCHRONIZING MULTIPLE GOROUTINE

// func main() {
// 	numGoroutine := 3
// 	done := make(chan int, 3)

// 	//create a goroutine inside the loop
// 	for i := range numGoroutine {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(time.Second)
// 			done <- id // sending signal of completion
// 		}(i)
// 	}

// 	for range numGoroutine {
// 	<-done // Wait for each goroutine to finish
// 	}

// 	fmt.Println("All goroutines are completed")
// }

//=========== SYNCHRONIZING DATA EXCHANGE

func main() {
	//An error will occurred during the running time.
	//It is simply because this channel(data := make(chan string)) is
	//non synchronize channel. To avoid the error, we need to close the channel(Close(data)) in the goroutine
	data := make(chan string)
	go func() {
		for i := range 5 {
			data <- "hello " + string(rune('0'+i))
			time.Sleep(100 * time.Millisecond)
		}
		//closing the channel
		close(data) //
	}()
	// Close(data) //Channel closed before the goroutine could send a value to the channel

	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	} // Loops over only on active channel, creates receiver each time and stops creating receiver(looping)
	// once is closed 
}

package main

import (
	"fmt"
	"time"
)

func main() {
  // == NON BLOICKING RECEIVE OPERATION
//  ch := make(chan int)
//
//  select{
//  case msg := <- ch:
//    fmt.Println("Receive:", msg)
//  default:
//    fmt.Println("No message aviable.")
//}
//
//  // == NON BLOCKING SEND OPERATION
//
//  select{
//  case ch <- 1:
//  fmt.Println("Sent message.")
//  default:
//  fmt.Println("Channel is not ready to receive")

//}


  // == NON BLOCKING OPERATION IN REAL TIME SYSTEM

  data := make(chan int)
  quit := make(chan bool)

  go func(){
for{
      select{
      case d := <- data:
        fmt.Println("Data received:", d)
      case <- quit:
      fmt.Println("Stopping...")
      return
        default:
        fmt.Println("Waiting for data...")
      time.Sleep(500 * time.Millisecond)
    }
    }
  }()

  for i := range 5{
    data <- i
    time.Sleep(1 * time.Second)
  }
}

package main

import (
	"fmt"
	"time"
)

func main(){
  timer1 := time.NewTimer(1 * time.Second)
  timer2 := time.NewTimer(2 * time.Second)

  select{
  case <- timer1.C:
  fmt.Println("Timer1 expired")
  case <- timer2.C:
  fmt.Println("Timer expired")
  }
}


//func main(){
//  timer := time.NewTimer(2 *time.Second)
//
//
//  go func(){
//    <- timer.C
//    fmt.Println("Delayed operations executed")
//  }()
//  fmt.Println("Waiting...")
//  time.Sleep(3 *time.Second) // this is a blocking timer starts
//  fmt.Println("End of program")
//}



//func longRunningOperation(){
//  for i := range 20{
//    fmt.Println(i)
//    time.Sleep(time.Second)
//  }
//}
//
//func main(){
//  timeout := time.After(2 * time.Second)
//  done :=make(chan bool)
//
//  go func(){
//    longRunningOperation()
//    done <- true
//  }()
//  select{
//  case <- timeout:
//  fmt.Println("Operation times out")
//  case <-done:
//  fmt.Println("Operation Completed")
//}
//
//}



//func main() {
//  time.Sleep(time.Second)
//  fmt.Println("Starting app.")
//  timer := time.NewTimer(2 * time.Second)
//  fmt.Println("Waiting for timer.c")
//  stopped := timer.Stop()
//  if stopped{
//    fmt.Println("timer Stopped")
//  }
//  timer.Reset(time.Second)
//  fmt.Println("Timer reset")
//  <- timer.C // blockking in nature
//  fmt.Println("Timer expired")
//}

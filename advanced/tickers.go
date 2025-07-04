package main

import (
	"fmt"
	"time"
)

func main(){
  ticker := time.NewTicker(time.Second)
  stop := time.After(5 * time.Second)

  defer ticker.Stop()


  for{

    select{
    case tick := <- ticker.C:
      fmt.Println("Tick at:", tick)
      case <- stop:
      fmt.Println("Stopping ticker.")
    return
  }
  }
}


//func periodicTask(){
//  fmt.Println("Performing task at:", time.Now())
//}

//func main(){
//  ticker := time.NewTicker(time.Second)
//  defer ticker.Stop()
//
//  for{
//    select {
//    case <- ticker.C:
//    periodicTask()
//  }
//  }
//}

 
//func main() { 
//  ticker := time.NewTicker(time.Second)
//
//  defer ticker.Stop()
//
 // for tick := range ticker.C{
 //   fmt.Println("Tick at:", tick)
 //   }


  //i := 1
  //for range 5{
  //  i  *= 2
  //  fmt.Println(i)
  //}
//}

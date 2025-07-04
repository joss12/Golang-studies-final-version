package main

import (
	"fmt"
	"sync"
)


func main(){
  var counter int
  var wg sync.WaitGroup
  var mu sync.Mutex


  numGoroutines := 5
  wg.Add(numGoroutines)

  increment := func(){
    defer wg.Done()
    for range 1000{
      mu.Lock()
      counter++
     mu.Unlock()
    }
  }
  for range numGoroutines{
    go increment()
  }
  wg.Wait()
  fmt.Printf("Final counter value: %d\n", counter)
}



//type Counter struct{
//  mu sync.Mutex
//  count int
//}
//
//func (c *Counter)increment(){
//  c.mu.Lock()
//  defer c.mu.Unlock()
//  c.count++
//}
//
//func (c *Counter)getValue()int{
//  c.mu.Lock()
//  defer c.mu.Unlock()
//  return c.count
//}
//
//func main() {
//  var wg sync.WaitGroup
//  counter := &Counter{}
//
//  numGoroutines := 10
//  
//  //wg.Add(numGoroutines)
//  for range numGoroutines {
//    wg.Add(1)
//    go func(){
//      defer wg.Done()
//      for range  1000{
//       counter.increment()
//        //counter.count++
//      }
//    }()
//  }
//
//  wg.Wait()
//  fmt.Println("Final counter value: %d\n", counter.getValue())
//}

package main

import (
	"sync"
	"sync/atomic"
  "fmt"
)

type AtomiCounter struct{
  count int64
}

func (ac *AtomiCounter)increment(){
  atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomiCounter)getValue() int64{
  return atomic.LoadInt64(&ac.count)
}

func main() {
  var wg sync.WaitGroup
  numgoroutines := 10
  counter := &AtomiCounter{}
  value := 0
    for range numgoroutines{
        wg.Add(1)
    go func(){
      defer wg.Done()
      for range 1000{
        counter.increment()
        value++
      }
    }()
  }

  wg.Wait()
  fmt.Printf("Final counter value %d\n", counter.getValue())
  //fmt.Printf("Final counter value %d\n", value)

}

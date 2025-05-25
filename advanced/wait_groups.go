package advanced

import (
	"fmt"
	"sync"
	"time"
)


// Construction example
type Worker struct{
  ID int
  Task string
}


func (w *Worker)PerformTask(wg *sync.WaitGroup){
defer wg.Done()
  fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
  time.Sleep(time.Second)
  fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func main(){
  var wg sync.WaitGroup
  //Define tasks to be performed by workers
  tasks := []string{"digging", "coding", "praying"}

  for i, task := range tasks{
    worker := Worker{ID: i+1, Task:task}
    wg.Add(1)
    go worker.PerformTask(&wg)
  }

  //wait for all workers to finish
  wg.Wait()

  //Contruction is finish
  fmt.Println("Contruction finish")
}


// EXAMPLE WITH CHANNELS
//func worker(id int, tasks <- chan int, results chan <- int, wg *sync.WaitGroup){
//  defer wg.Done()
//  fmt.Printf("WorkerID %d starting.\n", id)
//  time.Sleep(time.Second) // similate some worker
//  for task := range tasks{
//    results <- task * 2
//  }
//  results <- id *2
//  fmt.Printf("WorkerID %d finished\n", id)
//
//}
//
//func main(){
//  var wg sync.WaitGroup
//  numWorkers := 3
//  numJobs := 5
//  results := make(chan int, numJobs)
//  tasks := make(chan int, numJobs)
//
//
//  wg.Add(numJobs)
//
//  for i := range numWorkers{
//    go worker(i+1, tasks, results, &wg)
//  }
//  
//  for i := range numJobs{
//    tasks <- i+1
//  }
//  close(tasks)
//
//  go func(){
//    wg.Wait()
//    close(results)
//  }() 
//
//  for result := range results{
//    fmt.Println("Result:", result)
//  }
//}
//

// === Bsic example
//func  worker(id int, wg *sync.WaitGroup){
//
//  defer wg.Done()
//  //wg.Add(1) Wrong practice
//
//  fmt.Printf("Worker %d starting\n", id)
//  time.Sleep(time.Second) // similate some time spent on processing the task
//  fmt.Printf("worker %d finished\n", id)
//}
//
//
//
//func main() {
//  var wg sync.WaitGroup
//  numWorkers := 3
//
//  wg.Add(numWorkers) //The correct way
//
//  // Launch workers
//  for i := range numWorkers{
//    go worker(i, &wg)
//  }
//
//  wg.Wait()
//  fmt.Println("All workers finished!")
//}

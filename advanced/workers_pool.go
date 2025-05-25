package advaned

import (
  "fmt" 
  "time"
)

type ticketRequest struct{
  personID int
  numTickets int
  cost int
}

//Similate processing of ticket requests
func ticketProcessor(request <- chan ticketRequest, result chan <- int){
  for req := range request{
    fmt.Println("processing %d ticket(s) of personID %d with total cost %d\n", 
      req.numTickets, req.personID, req.cost)
    //similate process in time
    time.Sleep(time.Second)
    result <- req.personID
  }
}


func main(){
  numRequests := 5
  price := 5
  ticketRequests := make(chan ticketRequest, numRequests)
  ticketResults := make(chan int)


  //start ticket processor/worker
  for range 3 {
    go ticketProcessor(ticketRequests, ticketResults)
  }

  //Send ticket requests
  for i := range numRequests{
  ticketRequests  <-ticketRequest{personID: i+1,
  numTickets: (i+1) * 2, cost: (i+1) * price}

  }
  close(ticketRequests)

  for range numRequests{
  fmt.Printf("Ticket for personID %d processed successfully!\n", <-ticketResults)
  }
}
 


//func worker(id int, tasks <-chan int, results chan <- int){
//  for task := range tasks{
//    fmt.Printf("Worker %d processing tak %d\n", id, task)
//    //similate some work
//    time.Sleep(time.Second)
//    results <- task * 2
//  }
//}
//
//
//func main() {
//  numWorkers := 3
//  numJobs := 10
//  tasks := make(chan int, numJobs)
//  results := make(chan int, numJobs)
//
//  //create workers
//
//  for i := range numWorkers{
//    go worker(i, tasks, results)
//  }
//
//  //Send values to the tasks channel
//  for i := range numJobs{
//    tasks <- i
//    //colllect the resultsl
//  }
//    close(tasks)
//
//  for range numJobs{
//    result := <- results
//    fmt.Println("Result:", result)
//  }
//}

package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWork(ctx context.Context){
  for {
    select{
    case <- ctx.Done():
    fmt.Println("Work cancelled", ctx.Err())
    return
    default:
    fmt.Println("Working...")
  }
    time.Sleep(500 * time.Millisecond)
  }
}

func main(){
  ctx := context.Background()
  //ctx, cancel := context.WithTimeout(ctx,2 *time.Second)
  ctx, cancel := context.WithCancel(ctx)
  go func(){
    time.Sleep(2 * time.Second) // siimulatin a heavy task. Consider this a heavy time consuming
   cancel()
  }()

  ctx = context.WithValue(ctx, "requestID", "id134312")
  ctx = context.WithValue(ctx, "name", "Eddy")
  ctx = context.WithValue(ctx, "IP", "dsd.34.3433.34")
  ctx = context.WithValue(ctx, "OS", "Linux")

  go doWork(ctx)
  time.Sleep(3 *time.Second)
  
  requestID := ctx.Value("requestID")
  if requestID != nil{
    fmt.Println("Request ID:", requestID)
  }else{
    fmt.Println("No request ID found.")
  }
  logWithcontext(ctx, "This is a test log message")
}


func logWithcontext(ctx context.Context, message string){
  requestIDVal := ctx.Value("requestID")
  log.Printf("RequestID %v - %v", requestIDVal, message)
}

//func checkEvenOdd(ctx context.Context, num int)string{
//  select{
//  case <- ctx.Done():
//  return "Operation canceled"
//
//  default:
//  if num%2 == 0{
//      return fmt.Sprintf("%d is even", num)
//    }else{
//      return fmt.Sprintf("%d is Odd", num)
//    }
//}
//}

//func main(){
//  ctx := context.TODO()
//
//  result := checkEvenOdd(ctx, 5)
//  fmt.Println("Result with Context.TODO():", result)
//
//  ctx = context.Background()
//  ctx, cancel := context.WithTimeout(ctx, 1* time.Second)
//  defer cancel()
//
//  result = checkEvenOdd(ctx, 10)
//  fmt.Println("Result from timeout context:", result)
//
//  time.Sleep(2 * time.Second)
//  result = checkEvenOdd(ctx, 15)
//  fmt.Println("Result after timeout:", result)
//
//}

//func main() {
//  todoContext := context.TODO()
//  contextBkg := context.Background()
//
//
//  ctx := context.WithValue(todoContext, "name", "Eddy")
//  fmt.Println(ctx)
//  fmt.Println(ctx.Value("name"))
//
//
//  ctx1 := context.WithValue(contextBkg, "city", "Seoul")
//  fmt.Println(ctx1)
//  fmt.Println(ctx1.Value("city"))
//
//}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
  pid := os.Getpid()
  fmt.Println("Process ID:", pid)
  sigs := make(chan os.Signal, 1)
  done := make(chan bool, 1)

  //Notify channel on interrup or terminate signals
  signal.Notify(sigs, syscall.SIGINT,  syscall.SIGUSR1, syscall.SIGHUP)
  go func(){
    sig := <- sigs
    fmt.Println("Received signal:", sig)
    //os.Exit(0)
    done <- true
  }()

  go func(){

    for {
      select{
      case <- done:
      fmt.Println("Stopping work due to signal")
      return
      default:
      fmt.Println("Working...")
      time.Sleep(time.Second)
    }
    }



    //sig := <- sigs

   // for sig:= range sigs{
   // switch sig{
   // case syscall.SIGINT:
   // fmt.Println("Receive SIGINT(interrup)")
   //// case syscall.SIGTERM:
   //// fmt.Println("Receive SIGTERM (Terminate)")
   // case syscall.SIGHUP:
   // fmt.Println("Receive SIGHUP (Hangup)")
   // case syscall.SIGUSR1:
   // fmt.Println("Receive SIGUSR1 (User defined signal 1)")
   // fmt.Println("User define function is executed")
   //   continue
   // }
   // fmt.Println("Graceful exit.")
   // os.Exit(0)

  //}

  }()

  //similate some work
  //fmt.Println("Working...")
  for {
  time.Sleep(time.Second)
  }
}


//tasklist - List of all processes on windows
// taskkill /F /PID <PID> :taskkill /F /PID 1234
//Stop-Process -Id 1234 - Force

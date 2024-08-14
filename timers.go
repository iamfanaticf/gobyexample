package main

import (
  "fmt"
  "time"
)

func main() {
  timer1 := time.NewTimer(time.Second * 1)
  <-timer1.C
  fmt.Println("this code should run after 1 sec")

  timer2 := time.NewTimer(time.Second * 2)
  go func() {
    <-timer2.C
    fmt.Println("this code will 2 sec after the initial 3 seconds")
  }()

  fmt.Println("some code is arriving")
  time.Sleep(time.Second * 2)

  timer3 := time.NewTimer(time.Second * 4)
  
  go func() {
    <- timer3.C
    fmt.Println("this code should fire after 4 more seconds but i will cancel it")
  }()

  stop := timer3.Stop()
  if stop {
    fmt.Println("some future code stopped instead")
  }

  
  
}

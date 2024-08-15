package main

import (
  "fmt"
  "time"
)

func main() {
  ticker := time.NewTicker(time.Second)
  done := make(chan bool)

  go func() {
    for {
    select {
    case <- done:
      return
    case t := <- ticker.C:
      fmt.Println(t)
    }
  }
  }()

  time.Sleep(time.Second * 4)
  done <- true
  ticker.Stop() 
  time.Sleep(time.Second * 9)
}

package main

import (
  "time"
  "fmt"
)

func main() {
  ticker := time.NewTicker(time.Second)
  done := make(chan bool)
  go func() {
  for{
    select {
    case t := <-ticker.C:
        fmt.Println(t)
    case <-done:
        fmt.Println("done")
        return
    }
  }
  }()
  time.Sleep(time.Second * 10)
  done <- true
}



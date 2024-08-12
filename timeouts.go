package main

import (
  "fmt"
  "time"
)

func main() {
  ch := make(chan string)

  go func() {
    time.Sleep(2 * time.Second)
    ch <- "value 2"
  }()

  select {
  case v:= <-ch:
    fmt.Println(" value received", v)
  case <- time.After(1 *time.Second):
      fmt.Println(" timeout 2")
  }

  go func() {
    time.Sleep(time.Second)
    ch <- "value 1"
  }()

  select {
  case v:= <- ch: 
    fmt.Println(" value received", v)
  case <- time.After(time.Second * 2): 
    fmt.Println(" timeout 1")
  }
  fmt.Println(<-ch) // intentionally used one chan to examine the behaviour 

}

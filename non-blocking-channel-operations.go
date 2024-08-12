package main

import "fmt"

func main() {
  messages := make(chan string)
  signals := make(chan string)

  select {
    case msg:= <- messages: 
    fmt.Println("message received ", msg)
  default:
    fmt.Println("no message received")
  }

  msg := "hi"
  select {
  case messages <- msg :
    fmt.Println("message sent", msg)
  default:
    fmt.Println("no message sent")
  }

  select {
  case msg := <- messages:
    fmt.Println("message received", msg)
  case sig := <- signals:
    fmt.Println("signal received", sig)
  default:
    fmt.Println("neither messages nor signals received")
  }
}

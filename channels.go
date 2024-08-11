package main

import "fmt"

func main() {
  msgChan := make(chan string)

  go func() { msgChan <- "ping"}()

  msg := <- msgChan

  fmt.Println(msg)
}

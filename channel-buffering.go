package main

import "fmt"

func main() {
  msgChan := make(chan string, 2)
  
  msgChan <- "buffered"
  msgChan <- "channel"

  fmt.Println(<-msgChan)
  fmt.Println(<-msgChan)
}

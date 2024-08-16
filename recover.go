package main

import "fmt"

func mayPanic() {
  panic("no prob prob")
}

func main() {
  defer func() {
    if r:=recover(); r != nil {
      fmt.Println("recovered. Error:\n", r)
    }
  }()
  mayPanic()
  fmt.Println("after may panic")
}

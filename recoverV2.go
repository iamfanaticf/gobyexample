package main

import "fmt"

func mayPanic() {
  panic("no prob prob")
}
func toCallmayPanic() {
  defer func() {
    if r:=recover(); r != nil {
      fmt.Println("recovered. Error:\n", r)
    }
  }()
  mayPanic()
}
func main() {
  toCallmayPanic()
  fmt.Println("after may panic")
}

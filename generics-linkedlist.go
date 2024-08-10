package main

import(
  "fmt"
)

type Node[T any] struct {
  value T
  next *Node[T]
}

type LL[T] struct {
  head *Node[T]
}

func (ll *LL[T]) push(val T) {
  if(head == nil) {
    head = Node{ T, nil}
    return
  }
  var temp = head
  for temp.next != nil {
    temp = temp.next
  }
  temp.next = Node{ T, nil}  
}

func main() {
  myLinkedList := LL[int]{}
  myLinkedList.push(1)
  myLinkedList.push(2)
  myLinkedList.push(3)
  myLinkedList.push(4)
}


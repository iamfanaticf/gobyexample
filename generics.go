package main

import "fmt"

func MapKeys[K comparable, V any] (m map[K]V) []K {
  r := make([]K, 0, len(m))
  for k := range m {
    r = append(r, k)
  }
  return r
}

func main() {
  m := make(map[string]int)
  m["a"] = 1
  m["b"] = 2
  m["c"] = 3
  fmt.Println(MapKeys(m))

  n := make(map[int]string)
  n[1] = "a"
  n[2] = "b"
  n[3] = "c"
  fmt.Println(MapKeys(n))
}

package main

import(
  "fmt"
  "slices"
)

func main() {
  strs := [] string {"c", "b", "a", "d"}
  slices.Sort(strs)
  fmt.Println(strs)

  fmt.Println(slices.IsSorted(strs))
}

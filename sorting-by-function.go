package main

import(
  "fmt"
  "cmp"
  "slices"
)

func main() {
  fruits := []string {"peach", "banana", "kiwi"}

  lenCmp := func(a, b string) int {
    return cmp.Compare(len(a), len(b))
  }

  slices.SortFunc(fruits, lenCmp)
  fmt.Println(fruits)
  
  ints := []int {1,2,3}
  slices.SortFunc(ints, func(a, b int) int {return b - a})
  fmt.Println(ints)

  type Person struct {
    name string
    age int
  }

  people := []Person {
    Person{name: "Jax", age: 37},
    Person{name: "TJ", age: 25},
    Person{name: "Alex", age: 72},
    Person{name: "Blex", age: 72},
  }

  slices.SortFunc(people,
  func(a, b Person) int {
    if (a.age != b.age) {
      return cmp.Compare(a.age, b.age)
    } else {
      return cmp.Compare(b.name, a.name)
    }
  })

  fmt.Println(people)
}

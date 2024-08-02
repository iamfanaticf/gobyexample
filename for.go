package main

import "fmt"

func main() {
	i := 1
	for i<=3{
		fmt.Printf("%v ",i)
		i++
	}
	fmt.Println()

	for j:=1; j<3; j++ {
		fmt.Printf("%v ",j)
	}
	fmt.Println()

	for i := range 3 {
		fmt.Printf("%v ",i)
	}
	fmt.Println()

	for {
		fmt.Println("infinite")
		break
	}
	
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("%v ",n)
	}
	var u uint8
	for u = range 256 {
		fmt.Printnln("this will not work as 256 can't be initialized to uint")
	}

}	
		
	

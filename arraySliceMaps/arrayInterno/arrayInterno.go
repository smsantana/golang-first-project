package main

import "fmt"

func main() {
	a1 := [4]int{1, 2, 3, 4}
	s1 := a1[:2]
	a1[1] = 5
	s1[0] = 6
	fmt.Println(a1, s1)
}

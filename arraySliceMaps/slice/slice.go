package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array tem um tamanho definido
	s1 := []int{1, 2, 3}  //slice tem um tamanho variavel
	fmt.Println(a1, s1)

	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2]

	fmt.Println(a2, s3)

	s4 := s2[:1]
	fmt.Println(s4)

	//slice tem um tamanho e um ponteiro para uma estrutura e array
}

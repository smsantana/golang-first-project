package main

import "fmt"

func main() {

	// define o tamanho do array pelos valores inseridos
	numeros := [...]int{1, 6, 3, 4, 5}

	// iterando qualquer elemento de um array
	for i, numero := range numeros {
		fmt.Printf("\n%v) %v", i, numero)
	}

	fmt.Println(" \n ")

	for _, numero := range numeros {
		fmt.Println(numero)
	}

}

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// conversao de tipos basicos
	x := 2.4
	y := 2

	// int para float
	fmt.Println(x / float64(y))

	//int para String

	// z := 123

	//convertendo de int para string
	fmt.Println("minha nota e:", strconv.Itoa(97))

	//convertendo de string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 120)

	//string to boolean
	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("falso")
	}

}

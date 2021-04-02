package main

import (
	"fmt"
	m "math"
)

/*
Primeiro Programa em GO
*/

func main() {
	fmt.Printf("Hello World %s\n", "Go")
	const PI float64 = 3.1415
	var raio = 3.2

	//forma reduzida de atribuir variavel
	area := PI * m.Pow(raio, 2)
	fmt.Println(area)

	//outras formas de definir variaveis
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	e, f := 5, 6

	println(e, f)
}

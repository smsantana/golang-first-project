package main

import "fmt"

func usandoCase(number float64) string {
	var n = int(number)
	switch n {
	case 1, 2:
		return "A"
	case 3:
		return "b"

	case 5:
		return "c"
	default:
		return "Numero Invalido"
	}
}

func usandoCaseSemExpress(number float64) string {
	var n = int(number)
	switch {
	case n > 1 && n <= 2:
		return "A"
	case n == 3:
		return "b"

	case n == 5:
		return "c"
	default:
		return "Numero Invalido"
	}
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "func"
	default:
		return "tipo nao conhecido"

	}

}

func main() {
	//case tradicional
	println(usandoCase(0.2))
	println(usandoCase(2))
	println(usandoCase(5))

	//case sem expressao
	println(usandoCaseSemExpress(0.2))
	println(usandoCaseSemExpress(2))
	println(usandoCaseSemExpress(5))

	// case por tipo de variavel
	fmt.Printf("\ntipo de parametro: %v", tipo("saul"))
	fmt.Printf("\ntipo de parametro: %v", tipo(12))
	fmt.Printf("\ntipo de parametro: %v", tipo(25.3))
	fmt.Printf("\ntipo de parametro: %v", tipo(true))
	fmt.Printf("\ntipo de parametro: %v", tipo(func() { println("ola") }))

}

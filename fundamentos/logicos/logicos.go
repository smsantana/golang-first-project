package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	compraTv52 := trab1 && trab2
	compraTv32 := trab1 != trab2
	compraSorvete := trab1 || trab2

	return compraTv52, compraTv32, compraSorvete
}

func main() {
	trab1, trab2 := true, true
	compraTv52, compraTv32, compraSorvete := compras(trab1, trab2)

	fmt.Printf("tv52:%v, tv32: %v, Sorvete: %v, saudavel:%v", compraTv52, compraTv32, compraSorvete, !compraSorvete)
}

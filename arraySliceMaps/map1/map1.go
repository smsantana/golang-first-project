package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[123] = "Saul Moreira"
	aprovados[456] = "Jesse Moreira"
	aprovados[789] = "Juvencio Moreira"

	// fmt.Println(len(aprovados), aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("\ncpf:%v nome: %s", cpf, nome)
	}

	delete(aprovados, 123)

	fmt.Println("\n", aprovados)

}

package main

import "fmt"

func main() {
	funcSalario := map[string]float64{
		"saul":  1000.0,
		"jesse": 2000.0,
	}

	for nome, salario := range funcSalario {
		fmt.Println(nome, salario)
	}

	delete(funcSalario, "saul")

	fmt.Println(funcSalario)
}

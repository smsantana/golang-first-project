package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"S": {"saul": 1000.0, "sabia souza": 3000.0},
		"J": {"jesse": 2000.0},
	}
	fmt.Println(funcPorLetra)
	for letra, funcionario := range funcPorLetra {
		for nome, salario := range funcionario {
			fmt.Printf("\nLetra:%v nome: %v salario: %f", letra, nome, salario)
		}

	}
}

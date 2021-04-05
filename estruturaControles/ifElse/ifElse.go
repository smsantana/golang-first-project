package main

import "fmt"

func resultNotas(nota float64) {
	if nota >= 7 {
		fmt.Println("Aluno acima da media com nota:", nota)
	} else if nota >= 6 {
		fmt.Println("Aluno na media com nota:", nota)
	} else {
		fmt.Println("Aluno reprovado com nota:", nota)
	}
}
func main() {

	resultNotas(7.0)
	resultNotas(6.0)
	resultNotas(5.0)

}

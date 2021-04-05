package main

import "fmt"

func main() {
	var notas [3]float64
	notas[0], notas[1], notas[2] = 7.0, 8.0, 9.0
	//println(notas[0], notas[1], notas[2])

	total := 0.0
	for i := 0; i < len(notas); i++ {
		fmt.Printf("nota: %2f\n", notas[i])
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("\nmedia:%2f", media)

}

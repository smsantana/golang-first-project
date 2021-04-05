package main

import (
	"fmt"
	"time"
)

func main() {
	c := 10

	i := 1

	//for mais simples
	for i <= 5 {
		fmt.Printf("%v\n", i)
		i++
	}

	// for mais completo tradicional
	for i := 0; i < c; i++ {
		println(i)
	}

	// for infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 5)
	}

}

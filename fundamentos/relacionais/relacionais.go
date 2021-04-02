package main

import (
	"fmt"
	"time"
)

type Pessoa struct {
	Nome string
}

func main() {
	a, b := "Banana", "Banana"
	fmt.Println("Strings", a == b)

	fmt.Println("!=", 1 != 2)
	fmt.Println("<>", 3 > 1)
	fmt.Println("<=>", 3 >= 2)

	d1, d2 := time.Unix(0, 0), time.Unix(0, 0)

	fmt.Println("datas:", d1 == d2)
	fmt.Println("datas:", d1.Equal(d2))

	fmt.Println(d1, " == ", d2)

	p1, p2 := Pessoa{"Joao"}, Pessoa{"Joao"}
	fmt.Println("Struct", p1 == p2)

}

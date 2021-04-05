package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	fmt.Println(numeroAleatorio())
	//if com init
	if i := numeroAleatorio(); i > 5 { // tambem suportado pelo switch
		println(i)
		println("Acima da media")
	} else {
		fmt.Println("Abaixo da media")
	}
}

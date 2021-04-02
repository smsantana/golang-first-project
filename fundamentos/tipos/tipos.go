package main

import (
	"fmt"
	"reflect"
)

func main() {

	// String delimitar apenas com aspas  duplas
	var name = "saul"
	fmt.Println("o tipo da var e: ", reflect.TypeOf(name))
	fmt.Println("o tamanho da String", name, " e ", len(name))
	s2 := `ola meu nome
	e
	saul
	e
	o
	seu ?`

	// Strings com quebra de linha usar  craze ``
	fmt.Println(s2)
	// int
	numberInt := 1234

	// float
	numberFloat := 1.2

	//boolean
	varBoolean := true

	//char nao existe ele convert para int32 e tem um mapeamento
	char := 'a'
	fmt.Println(char)

	fmt.Printf("String:%s Int:%v, Float: %v, Boolean: %v", name, numberInt, numberFloat, varBoolean)
}

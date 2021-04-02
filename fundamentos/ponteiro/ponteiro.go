package main

import "fmt"

func main() {
	/*
	   o esteriscos e a forma de declarar uma variavel como ponteiro,
	*/
	i := 1
	j := "saul"

	//declaracao de ponteiros
	var q *string = nil
	var p *int = nil

	q = &j           //criando referencia a variavel j
	*q += " moreira" // atribuindo valor via ponteiro a variavel j
	fmt.Println(j)   // imprimindo valor da variavel j apos concatenar string via ponteiro

	p = &i // pegando o endereco da variavel i
	*p++   // desreferenciando e incrementando(pegando o valor do ponteiro que aponta para var i)
	i++

	println(p, *p, i, &i)

}

package main

import "fmt"

//Ponteiro é uma referencia de memoria

func main() {
	var variavel int
	var ponteiro *int

	variavel = 100
	ponteiro = &variavel

	fmt.Println(variavel, ponteiro)

	fmt.Println(variavel, *ponteiro)

	fmt.Println(&variavel, *ponteiro)
}

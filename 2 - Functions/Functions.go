package main

import "fmt"

func Calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma, subtracao := Calculos(1, 2)
	fmt.Println(soma, subtracao)
}

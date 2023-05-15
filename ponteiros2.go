// Escreva uma função que receba um ponteiro para um inteiro e verifique se esse inteiro é par ou ímpar.
// A função deve atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar.
package main

import "fmt"

func ptrParImpar(p *int) {
	if *p%2 == 0 {
		*p = 0
	} else if *p%2 != 0 {
		*p = 1
	}
}

func main() {
	num := 95

	ptrParImpar(&num)
	fmt.Println(num)
}

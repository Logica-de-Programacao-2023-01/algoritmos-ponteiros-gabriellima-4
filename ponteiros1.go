// Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n.
// A função deve atualizar o valor do inteiro para a soma dos n primeiros números naturais.
package main

import "fmt"

func somarNaturais(p *int, n int) {
	soma := 0

	for i := 1; i <= n; i++ {
		soma += i
	}
	*p = soma
}

func main() {
	s := 0

	somarNaturais(&s, 10)
	fmt.Println(s)
}

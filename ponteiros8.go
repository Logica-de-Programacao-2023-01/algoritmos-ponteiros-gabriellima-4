// Crie uma função que receba um ponteiro para uma variável como argumento
// e modifique o valor da variável dentro da função.
// Certifique-se de que o ponteiro aponte para uma área de memória válida e que a memória seja liberada após o uso.
package main

import "fmt"

func modificarValor(v *int) {
	x := 0
	fmt.Println("Digite um valor para ser alterado.")
	fmt.Scanln(&x)

	*v = x
}

func main() {
	x := 0
	modificarValor(&x)
	fmt.Println("O novo valor agora é: ", x)
}

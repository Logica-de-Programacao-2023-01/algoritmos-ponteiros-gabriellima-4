// Escreva uma função em Go que receba um ponteiro para uma variável inteira
// e atualize o valor da variável para a soma dos valores dos seus dois últimos dígitos
// (por exemplo, se o valor inicial da variável for 1234, o novo valor será 3+4=7).
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func somaDoisUltimosDigitos(num *int) {
	numeroEmString := strconv.Itoa(*num)
	separador := strings.Split(numeroEmString, "")

	d1 := 0
	d2 := 0
	soma := 0

	i := len(separador) - 1

	d1, _ = strconv.Atoi(separador[i])
	d2, _ = strconv.Atoi(separador[i-1])
	soma = d1 + d2

	*num = soma

}

func main () {
	x := 28052022  // Data da Final da Champions 2022
	
	somaDoisUltimosDigitos(&x)
	fmt.Println(x)
}

// Escreva uma função em Go que receba um ponteiro para uma string e atualize o valor da string para seu reverso.
package main

import (
	"fmt"
	"strings"
)

func reverterString(str *string) {
	novaString := strings.Split(*str, "")
	*str = ""

	for i := len(novaString) - 1; i >= 0; i-- {
		*str += novaString[i]
	}
}

func main() {
	s := "tobiaS booN"

	reverterString(&s)
	fmt.Println(s)
}

// Escreva uma função em Go que receba um ponteiro para uma variável float64
// e atualize o valor da variável para a média aritmética entre o seu valor atual e o valor da constante Pi.
package main

import (
	"fmt"
	"math"
)

func ptrMediaFloat(v *float64) {
	var x float64
	fmt.Print("Digite um valor para x:")
	fmt.Scan(&x)

	x = (x - math.Pi) / 2
	*v = x

}

func main() {
	var x float64
	ptrMediaFloat(&x)
	fmt.Println(x)
}

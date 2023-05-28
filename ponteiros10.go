// Implemente uma função que receba um ponteiro para uma slice e seu tamanho
// e preencha-o com os n primeiros números primos.
package main

import "fmt"

func slicePrimos(sl *[]int, n int) {
	numPrimo := true

	for i := 2; len(*sl) != n; i++ {

		for k := 2; k < i; k++ {

			if i%k == 0 {
				numPrimo = false
				break
			}

		}
		if numPrimo == true {
			*sl = append(*sl, i)
		} else {
			numPrimo = true
		}
	}
}

func main() {
	numeros := []int{}

	slicePrimos(&numeros, 7)
	fmt.Println(numeros)
}

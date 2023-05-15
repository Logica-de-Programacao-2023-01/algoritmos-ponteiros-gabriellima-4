// Implemente uma função que receba um ponteiro para uma struct "Livro" com campos "Título", "Autor" e "Preço",
// e que adicione um desconto de 10% no preço do livro.
package main

import "fmt"

type Livro9 struct {
	Titulo string
	Autor  string
	Preco  float64
}

func bookDiscount(v *Livro9) {
	v.Preco = v.Preco - v.Preco*0.10
}

func main() {
	y := Livro9{
		Titulo: "Vivências de um taxista",
		Autor:  "Agostinho Carrara",
		Preco:  30.90,
	}
	fmt.Println("Antes do desconto:", y)

	bookDiscount(&y)
	fmt.Println("Depois do desconto, o preço do livro agora é:", y)
}

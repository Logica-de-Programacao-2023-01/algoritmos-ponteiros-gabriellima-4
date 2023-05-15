// Escreva uma função em Go que receba um ponteiro para um struct Livro com campos título e autor,
// e altere o título do livro para "Desconhecido" se o autor for "Anônimo".
package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func unknownWriter(l *Livro) {
	novoA := "Desconhecido"
	if l.Autor == "Anônimo" {
		l.Autor = novoA
	}
}

func main() {
	x := Livro{
		Titulo: "Vivências de João Andante",
		Autor:  "Anônimo",
	}
	fmt.Println("Antes da mudança:", x)

	unknownWriter(&x)
	fmt.Println("Depois da mudança:", x)
}

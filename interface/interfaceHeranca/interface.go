package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome string
	sobrenome string
}

type produto struct {
	valor float64
	nome string
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - %.2f", p.nome, p.valor)
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func imprimir(i imprimivel){
	fmt.Println(i.toString())
}

func main() {
	// Implementações de interface sao IMPLICITAS em GO, produto e pessoa implementam,nesse caso, a mesma interface (imprimivel)
	// Dessa forma, imprimivel consegue encapsular qualquer struct que implementa os metodos da interface
	// porem obedecendo o conceito de polimorfismo, de forma que produto e pessoa tem seus proprios "jeitos" quando se é chamada o toString()

	var qualquercoisa imprimivel = produto{nome : "PS4", valor: 2000}
	fmt.Println(qualquercoisa.toString())
	qualquercoisa = pessoa{"Thyago", "Freitas"}
	fmt.Println(qualquercoisa.toString())
	fmt.Println("/////////////////////////////////////")

	pessoaX := pessoa{"Thyago", "Silva"}
	imprimir(pessoaX)
}
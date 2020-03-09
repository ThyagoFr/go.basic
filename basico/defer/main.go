package main

import "fmt"

// O defer faz com que um trecho de código seja executado antes da chamada de um return
// Muito util para funceso que operam sobre conexoes em banco de dados
// Assim vc pode definir que a conexao com o banco,por exemplo, sempre seja executado

func imprimir(numero int) int {
	defer fmt.Println("Saindo...")
	if numero >= 5000 {
		fmt.Println("Maior ou igual a 5000")
		return numero
	}
	fmt.Println("Menor que 5000")
	return numero
}

func main() {
	fmt.Println(imprimir(4000))
}

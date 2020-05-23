package main

import (
	"fmt"
)

func multiplicar(valor1 float64, valor2 float64) float64 {
	return valor1 * valor2
}

func soma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func exec(funcao func(op1, op2 float64) float64, op3, op4 float64) float64 {
	return funcao(op3, op4)
}

func main() {

	/*
		y, err1 := strconv.ParseFloat(os.Args[1], 64)
		x, err2 := strconv.ParseFloat(os.Args[2], 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Erro na conversão")
		} else {
			fmt.Printf("%.2f * %.2f = %.2f \n", x, y, multiplicar(x, y))
			fmt.Printf("%.2f + %.2f = %.2f \n", x, y, soma(x, y))
		}
	*/

	// Salvando funcao em uma variavel
	funcSub := func(op1, op2 float64) float64 {
		return op1 - op2
	}

	// Passando uma funcao como parametro

	fmt.Println(funcSub(1.5, 2.4))
	fmt.Println(exec(soma, 20.0, 1.5))
	fmt.Println(exec(funcSub, 1.14, 2.90))
}

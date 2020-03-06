package main

import (
	"fmt"
	"os"
	"strconv"
)

func multiplicar(valor1 float64, valor2 float64) float64 {
	return valor1 * valor2
}

func soma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func main() {
	y, err1 := strconv.ParseFloat(os.Args[1], 64)
	x, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Erro na conversão")
	} else {
		fmt.Printf("%.2f * %.2f = %.2f \n", x, y, multiplicar(x, y))
		fmt.Printf("%.2f + %.2f = %.2f \n", x, y, soma(x, y))
	}
}

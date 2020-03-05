package main

import (
	"fmt"
	"os"
	"strconv"
)

func buscaBin(numeros []float64, n float64) bool {
	esquerda, direita := 0, len(numeros)-1
	for esquerda <= direita {
		meio := int((esquerda + direita) / 2)
		if numeros[meio] == n {
			return true
		} else if numeros[meio] > n {
			direita = meio - 1
		} else {
			esquerda = meio + 1
		}
	}
	return false
}

func main() {
	sliceY := os.Args[2:]
	input, err := strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("Input = %.2f \n", input)
	if err != nil {
		fmt.Println("Valor de busca invalida")
	}
	sliceZ := make([]float64, len(sliceY))
	for index, n := range sliceY {
		value, err := strconv.ParseFloat(n, 64)
		if err != nil {
			fmt.Println("Houve um erro")
		} else {
			sliceZ[index] = value
		}
	}
	fmt.Println(sliceZ)
	fmt.Println(buscaBin(sliceZ, input))
}

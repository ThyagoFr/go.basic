package main

import (
	"fmt"
	"os"
	"strconv"
)

func quicksort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)
	indicePivo := len(n) / 2
	pivo := n[indicePivo]
	n = append(n[:indicePivo], n[indicePivo+1:]...)
	menores, maiores := particionar(n, pivo)
	return append(
		append(quicksort(menores), pivo),
		quicksort(maiores)...)
}

func particionar(
	numeros []int,
	pivo int) (menores, maiores []int) {

	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}
	return menores, maiores
}

func main() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))
	for index, n := range entrada {
		valor, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Houve erro")
		} else {
			numeros[index] = valor
		}
	}
	fmt.Println(quicksort(numeros))
}

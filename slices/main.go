package main

import (
	"fmt"
)

func multiplica(array []int) {
	for index, v := range array {
		array[index] = v + 2
	}
}

func teste(x [5]int) [5]int {
	for index, v := range x {
		x[index] = v * 2
	}
	return x
}

func main() {
	slice := make([]int, 10)
	multiplica(slice)
	fmt.Println(slice)
	array := [...]int{1, 2, 3, 4, 5}
	arrayRetorno := teste(array)
	fmt.Println(arrayRetorno)

	// Fatiando slices
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice3 := slice2[:5] // 5 é um limite aberto(só vai ate o indice 4)
	fmt.Println("Slice inteiro")
	fmt.Println(slice2)
	fmt.Println("Slice fatiado da posicao 0 à 4")
	fmt.Println(slice3)

	// Inserindo coisas no slice (ultima posicao)
	fmt.Println(slice2)
	slice2 = append(slice2, 50)
	fmt.Println(slice2)

	// Inserindo valor no comeco do slice
	fmt.Println(slice2)
	slice4 := []int{0, 0, 0}
	slice2 = append(slice4, slice2...)
	fmt.Println(slice2)

	// Inserindo em qualquer posicao
	slice5 := []int{10, 11, 12, 17, 18, 19}
	slice6 := []int{13, 14, 15, 16}
	slice5 = append(slice5[:3], append(slice6, slice5[3:]...)...)
	fmt.Println(slice5)

	// Copiando slices
	slice7 := make([]int, len(slice6))
	copy(slice7, slice6)
	for index := range slice7 {
		slice7[index] *= 4
	}
	fmt.Println(slice6)
	fmt.Println(slice7)
}

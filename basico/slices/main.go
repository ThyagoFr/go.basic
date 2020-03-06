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
}

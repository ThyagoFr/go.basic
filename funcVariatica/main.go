package main

import "fmt"

func imprimir(numeros ...float64) {
	for _, value := range numeros {
		fmt.Println(value)
	}
}
func main() {
	imprimir(1.2, 3.2, 5.4, 10)
	valores := []float64{1.2, 4, 5, 10.123}
	fmt.Println("Imprimindo slices...")
	imprimir(valores...)
}

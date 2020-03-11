package main

import "fmt"

func main() {
	// 10 AMOSTRAS
	valores := []float64{100,105,115,130,150,175,205,240,280,325,375,430,490,555,625,700,780,865,955,1050}
	n := 10
	valor := 0.0
	for j := n; j<len(valores); j++ {
		fmt.Println(j,valores[j])
		for i:= 0; i<=10; i++{
			valor += valores[j-i]
		}
		valores[j] = valor/float64(n)
	}
	fmt.Println(valores)
	fmt.Println(len(valores))
}

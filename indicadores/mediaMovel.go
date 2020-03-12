package main

import (
	"fmt"
)

func main() {
	// 10 AMOSTRAS
	media := []float64{}
	valores := []float64{100,105,115,130,150,175,205,240,280,325,375,430,490,555,625,700,780,865,955,1050}
	n := 10
	for j := n-1; j<len(valores); j++ {
		valor := 0.0
		for i:= 0; i<n; i++{
			valor += valores[j-i]
		}
		k := valor/float64(n)
		media = append(media,k)
	}
	fmt.Println(media)
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	valores := []float64{180.0,280.0,190.0,240.0}
	medias := []float64{}
	periodo := os.Args[1]
	periodoInt,_ := strconv.Atoi(periodo)
	pesos := os.Args[2:]
	pesosFloat := []float64{}

	somaPesos := 0.0
	for _,value := range pesos {
		v,_ := strconv.ParseFloat(value,64)
		somaPesos += v
		pesosFloat = append(pesosFloat,v)
	}

	sum := 0.0
	for i := 0; i < periodoInt; i++ {
		sum += valores[i]* pesosFloat[i]
	}

	fmt.Println(sum)
	media := sum / somaPesos
	medias = append(medias, valores[:periodoInt]...)
	medias = append(medias, media)
	for j := periodoInt; j < len(valores); j++ {
		media = (media*somaPesos - valores[j-periodoInt]*pesosFloat[0] + valores[j]*pesosFloat[periodoInt-1]) / somaPesos
		medias = append(medias, media)
	}

	fmt.Println(valores,medias,pesos,periodo)
}

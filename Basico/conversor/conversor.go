package main

// Argumentos passados por linhas de comando
// sao sempre strings
// Esses argumentos estao em os.Args

import (
	"fmt"     // Biblioteca para saida (printf)
	"os"      // Lida com servicos relacionados ao S.O
	"strconv" // Conversao de strings para outros tipos
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso : conversor <valores> <unidade>")
		os.Exit(1)
	} else {
		var unidadeOrigem = os.Args[len(os.Args)-1]
		valoresOrigem := os.Args[1 : len(os.Args)-1] // ou unidadeOrigem := os.Args[1] - Forma resumida com inferencia de tipo
		var unidadeDestino string
		if unidadeOrigem == "celsius" {
			unidadeDestino = "fahrenheit"
		} else if unidadeOrigem == "quilometros" {
			unidadeDestino = "milhas"
			for index, value := range valoresOrigem {
				valorOrigem, err := strconv.ParseFloat(value, 64)
				if err != nil {
					fmt.Println("Um erro ocorreu ao converter um dos valores de entrada")
				} else {
					fmt.Println(index, valorOrigem)
				}
			}
		} else {
			fmt.Printf("%s nao é uma unidade conhecida", unidadeDestino)
			os.Exit(1)
		}
		fmt.Printf("Unidade Origem : %s \n", unidadeOrigem)

	}
}

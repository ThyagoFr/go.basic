package main

import (
	"fmt"
	"os"
	"strings"
)

type Estado struct {
	nome      string
	populacao int
	capital   string
}

func colherEst(palavras []string) map[string]int {
	est := make(map[string]int) // est := map[string]int{}
	for _, palavra := range palavras {
		inicial := strings.ToUpper(string(palavra[0]))
		contador, encontrado := est[inicial]
		// Retorna 2 valores, o primeiro é o valor associado a chave
		// O segundo é um bool pra saber se a chave existe
		if encontrado {
			est[inicial] = contador + 1
		} else {
			est[inicial] = 1
		}

	}
	return est
}

func main() {
	palavras := os.Args[1:]
	estatisticas := colherEst(palavras)
	for inicial, contador := range estatisticas {
		fmt.Printf("%s = %d \n", inicial, contador)
	}

	estados := map[string]Estado{}
	estados["GO"] = Estado{"Goias", 000001, "Goiania"}
	fmt.Println(estados)

	// Deletando valor do map
	fmt.Println("Apagando a chave GO...")
	delete(estados, "GO")
	fmt.Println(estados)

	// Checando se a chave existe
	_, existe := estados["GO"]
	if existe {
		fmt.Println("A chave GO nao existe")
	} else {
		fmt.Println("A chave nao existe")
	}
	// Ordem em MAPS nao é garantida
	estados["GO"] = Estado{"Goias", 000001, "Goiania"}
	estados["CE"] = Estado{"Ceara", 000001, "Fortaleza"}
	// Iterando sobre maps
	for chave, valor := range estados {
		fmt.Printf("%s (%s) possui %d habitantes. \n", valor.nome, chave, valor.populacao)
	}
}

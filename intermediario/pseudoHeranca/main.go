package main

import "fmt"

type carro struct {
	portas int
	modelo string
	valor  float64
}

type ferrari struct {
	carro      // Campos anonimos para simular uma herança
	velocidade float64
}

func main() {
	ferrari1 := ferrari{}
	ferrari1.portas = 4
	ferrari1.modelo = "seila"
	ferrari1.valor = 10000
	ferrari1.velocidade = 100
	fmt.Println(ferrari1)
}

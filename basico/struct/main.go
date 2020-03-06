package main

import (
	"errors"
	"fmt"
)

type Pilha struct {
	valores []interface{}
}

func (p Pilha) Tamanho() int {
	return len(p.valores)
}

// Sintaxe para associar metodos a structs
func (p Pilha) Vazia() bool {
	return p.Tamanho() == 0
}

// Trabalha em cima do ponteiro que possui o endereco para a pilha
func (p *Pilha) Empilhar(valor interface{}) {
	p.valores = append(p.valores, valor)
}

// Trabalha em cima do ponteiro que possui o endereco para a pilha
func (p *Pilha) Desempilhar() (interface{}, error) {
	if p.Vazia() {
		return nil, errors.New("Pilha Vazia")
	}
	valor := p.valores[p.Tamanho()-1]
	p.valores = p.valores[:p.Tamanho()-1]
	return valor, nil
}

func (p Pilha) Mostrar() {
	for _, value := range p.valores {
		fmt.Println(value)
	}
}

func main() {
	pilha := Pilha{}
	fmt.Println("Pilha criada com tamanho ", pilha.Tamanho())
	fmt.Println("Pilha vazia? ", pilha.Vazia())
	pilha.Empilhar("Go1")
	pilha.Empilhar("Go2")
	pilha.Empilhar("Go3")
	pilha.Empilhar("Go4")
	pilha.Mostrar()
}

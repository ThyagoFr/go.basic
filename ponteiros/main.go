package main

import "fmt"

type Arquivo struct {
	nome     string
	path     string
	extensao string
}

func main() {
	arq := Arquivo{"texto", "src/bin/", "txt"}
	ponteiroArq := &arq
	fmt.Println(arq)
	// Operador & retorna um ponteiro para a struct criada
	fmt.Println(ponteiroArq)

	listaArq := []Arquivo{
		{"arquivo1", "src/bin/", "txt"},
		{"arquivo2", "src/bin/", "txt"},
		{"arquivo3", "src/bin/", "txt"}}
	plistaArq := &listaArq
	// Nao é possivel fazer operacoes aritmeticas (soma e afins com ponteiros)
	// Apenas comparaçoes
	// *ponteiro da acesso ao valor do armazenado na posicao de memoria para qual o ponteiro aponta

	fmt.Println(*plistaArq)
}

package main

import "fmt"

func main() {
	// Loops nomeados
externo:
	for {
		for i := 0; i < 10; i++ {
			if i == 5 {
				fmt.Println("Saindo do loop externo")
				break externo
			}
		}
	}
	// Implementacao de um while
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}

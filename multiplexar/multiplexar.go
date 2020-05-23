package main
// Multiplexar - Inserir valores de varios canais em um canal

import (
	"fmt"
	"meupacote/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <- origem // Enquanto o canal de origem tiver dados, os dados serao repassados para o canal de destino
		// Quando nao tiver dados, o for irá "parar"
	}
}

func juntar (entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1,c)
	go encaminhar(entrada2,c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br","https://www.google.com"),
		html.Titulo("https://www.youtube.com", "https://www.apple.com"),
		)
	// Será printado o primeiro titulo dos 4 (o q foi obtido primeiro)
	fmt.Println(<-c)
}
package main

import (
	"fmt"
	"time"
)
// Canais com Buffer
// A operacao de insercao de valores em canal só trava
// quando vc encher o buffer (nesse caso 3 valores)
// Quando for inserir o quarto valor ele ira travar, até
// que alguem leia um valor do buffer
func inserir(c chan int) {
	for i := 1; i <= 6; i++ {
		c <- i
		fmt.Printf("Enviou %d \n",i)
	}
}
func main(){
	ch := make(chan int, 3)
	go inserir(ch)
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}

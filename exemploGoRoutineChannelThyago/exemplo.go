package main

import (
	"fmt"
	"time"
)

func inserir(c chan int){
	for i := 1; i <=10; i++ {
		c <-i // insere e trava esperando a leitura
		// depois de lido, espera 2 segundos
		time.Sleep(time.Second * 2)
	}
}
func main(){
	c := make(chan int)
	go inserir(c)
	for i := 1; i <= 10; i++ {
		// trava esperando alguem inserir um valor
		fmt.Println(<-c)
	}
}

package main

import (
	"fmt"
)

func add(ch chan int) {
	fmt.Println("Adicionando...")
	ch<-1
	fmt.Println("Adicionado")
	ch<-2
	fmt.Println(".........")
}

func main() {
	ch := make(chan int)
	go add(ch)
	fmt.Println("Hm....")
	fmt.Println(<-ch)

	// A opercao de leitura de um canal trava a execucao do processo
	// ate que a leitura tenha sido feita(apos uma goroutine ter adicionado ao canal).
}
package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operacao bloquenate, só é possivel adicionar um novo valor,
	// quando alguem ler o conteudo do canal (caracteristica especifica de um canal sem buffer)
	fmt.Println("Só depois que foi lido...")
}

func main(){
	c := make(chan int) // canal sem buffer
	go rotina(c)
	fmt.Println(<-c) // Só passa dessa linha quando tiver algo no canal a ser lido.
	fmt.Println(<-c) // Causa de um deadlock pois todas as go routines ja morreram
	// e o programa ficaria esperando algo que nao vai acontecer (uma go routine adicionar algo no canal)
	fmt.Println("FIM (NUNCA SERA VISTO")
}

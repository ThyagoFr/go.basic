package main

import (
	"fmt"
	"time"
)

// Canal é a forma de comunicacao entre goroutines
// Canal é um tipo da linguagem

func doisTresQuatroVezes (base int, c chan int){
	fmt.Println("Inserindo primeiro numero")
	time.Sleep(time.Second)
	c <- 2* base
	time.Sleep(time.Second)

	// Essa operacao só sera realizada quando o valor inserido
	// anteriormente (2*base) for retirado do canal
	fmt.Println("Inserindo segundo numero")
	c <- 3* base
	time.Sleep(3*time.Second)
	fmt.Println("Inserindo terceiro numero")
	c <- 4* base

}
func main() {
	// Canal sem buffer
	c := make(chan int)
	go doisTresQuatroVezes(2,c)
	fmt.Println("A")
	a := <-c
	b := <-c
	fmt.Println("B")
	fmt.Println(a,b)
	fmt.Println(<-c)
}

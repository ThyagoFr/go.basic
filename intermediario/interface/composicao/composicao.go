package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}
type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw7 struct {}

func (b bmw7) fazerBaliza(){
	fmt.Println("Baliza...")
}

func (b bmw7) ligarTurbo(){
	fmt.Println("Turbo...")
}

func main() {
	var coisa1 luxuoso = bmw7{}
	coisa1.fazerBaliza()
	var coisa2 esportivo = bmw7{}
	coisa2.ligarTurbo()
	var coisa3 esportivoLuxuoso = bmw7{}
	coisa3.fazerBaliza()
}
package main

import (
  "fmt"
)

func main() {
  // ... ellipsis (instrui o compilador a contar com base na quantidade de arrays declarados)
  // Arrays - Colexoes de mesmo tipo e tamanho fixo e invariavel
  numeros := [5]int{1, 2, 3, 4, 5}
  primos := [...]int{2, 3, 5, 7, 11, 13}
  var nomes [2]string // nomes := [2]strings{}
  var a [3]int
  fmt.Println(primos)
  fmt.Println(a)
  fmt.Println(numeros)
  fmt.Println(nomes)

  // matrizes
  matrizA := [2][2]int{{1, 2}, {3, 4}}
  fmt.Println(matrizA)
  for _, value := range matrizA {
    fmt.Println(value)
  }
}

package main

import (
	"fmt"
	"meupacote/html"
)

func main() {
	resposta := html.Titulo("https://www.google.com","https://www.google.com")
	fmt.Println(<-resposta,<-resposta)
}

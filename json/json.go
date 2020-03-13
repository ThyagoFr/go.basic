package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
	Tags []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := produto{1,"Notebook",2200.0,[]string{"Promocao","Eletrônico"}}
	p1Json,_ := json.Marshal(p1)
	fmt.Println(string(p1Json))
	// p1Json é na verdade um slice de bytes,entao é preciso converter para string
	// fmt.Println(p1Json)

	var p2 produto
	p2Json := `{"id":2,"nome":"Samsung X","preco":1200,"tags":["Promocao","Eletrônico"]}`
	json.Unmarshal([]byte(p2Json),&p2)
	fmt.Println(p2.Tags[1])


}

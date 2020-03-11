package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type usuario struct {
	id int
	nome string
}

func main(){
	db,err := sql.Open("mysql",  "root:root@tcp(172.17.0.2:3306)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows,_ := db.Query("SELECT * from usuarios")
	defer rows.Close()
	for rows.Next() {
		us := usuario{}
		rows.Scan(&us.id,&us.nome)
		fmt.Println(us)
	}

}
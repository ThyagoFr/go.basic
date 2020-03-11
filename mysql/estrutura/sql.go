package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main(){
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	exec(db, "CREATE DATABASE IF NOT EXISTS cursogo")
	fmt.Println("Criou o banco...")
	exec(db, "USE cursogo")
	fmt.Println("Mudando para o banco de dados criado")
	exec(db, `CREATE TABLE usuarios (
								id integer auto_increment,
								nome varchar(80),
								PRIMARY KEY (id)
	)`)
	fmt.Println("Criando tabela usuarios")
}
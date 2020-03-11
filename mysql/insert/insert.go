package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) values (?)")
	stmt.Exec("Maria")
	stmt.Exec("Joao")
	res,_ := stmt.Exec("Pedro")
	id, _ := res.LastInsertId()
	fmt.Println(id)
}
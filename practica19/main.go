package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.26.0.39:32769)/prueba")
	if err != nil {
		log.Fatalln("Error Abriendo base", err)
	}
	defer db.Close()

	if db.Ping() == nil {
		fmt.Println("La base esta abierta!!")
	} else {
		fmt.Println("No responde la conexión a la base")
	}
}

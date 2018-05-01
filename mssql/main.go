package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	condb, errdb := sql.Open("mssql", "server=oficina.xuitec.com;user id=externo;password=externo;database=lecContador")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	var (
		sqlversion string
	)
	rows, err := condb.Query("select @@version")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&sqlversion)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(sqlversion)
	}

	var valor int
	errr := condb.QueryRow("select count(*) from dev_data_inst").Scan(&valor)
	if errr != nil {
		log.Fatal("error : ", errr)
	}
	log.Println("-->", valor)

	defer condb.Close()
}

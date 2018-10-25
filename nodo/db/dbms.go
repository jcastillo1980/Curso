package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	hostDBMS     = "oficina.xuitec.com"
	portDBMS     = 1433
	userDBMS     = "sa"
	passwordDBMS = "cst2014C"
	nameBaseDBMS = "lecContador"
)

func init() {
	log.Println("Inicio DB MSSQL")

}

// ListaNombresMS ???
func ListaNombresMS() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", hostDBMS, userDBMS, passwordDBMS, portDBMS, nameBaseDBMS)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Panicln("ERROR ListaNombres():", err)
		return
	}

	defer conn.Close()

	rows, err := conn.Query("select id,nombre from dev_base")

	for rows.Next() {
		var nombre string
		var id int
		err = rows.Scan(&id, &nombre)
		if err != nil {
			log.Println("ERROR ListaNombresMS() Query:", err)
			return
		}
		log.Println("ID:", id, "NOMBRE:", nombre)
	}

}

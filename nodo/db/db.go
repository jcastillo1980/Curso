package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	log.Println("Inicio DB")
}

var (
	hostDB     = "oficina.xuitec.com"
	portDB     = 3306
	userDB     = "externo"
	passwordDB = "externo"
	nameBaseDB = "blx2"
)

func mu(a ...interface{}) []interface{} {
	return a
}

// ListaNombres ????
func ListaNombres() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userDB, passwordDB, hostDB, portDB, nameBaseDB))
	if err != nil {
		log.Println("ERROR ListaNombres():", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select id,nombre from dev_base")

	for rows.Next() {
		var nombre string
		var id int
		err = rows.Scan(&id, &nombre)
		if err != nil {
			log.Println("ERROR ListaNombres() Query:", err)
			return
		}
		log.Println("ID:", id, "NOMBRE:", nombre)
	}

}

// UpdateName ???
func UpdateName(id int, nombre string) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userDB, passwordDB, hostDB, portDB, nameBaseDB))
	if err != nil {
		log.Println("ERROR UpdateName():", err)
		return
	}
	defer db.Close()

	result, err := db.Exec("update dev_base set nombre = ? where id = ?", nombre, id)
	if err != nil {
		log.Println("ERROR UpdateName() Exec:", err)
		return
	}

	log.Println(":", mu(result.RowsAffected())[0].(int64))
}

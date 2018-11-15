package db

import (
	"database/sql"
	"fmt"
	"log"
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

// Setup ?????
func Setup(h string, p int, u string, pw string, n string) {
	hostDB = h
	portDB = p
	userDB = u
	passwordDB = pw
	nameBaseDB = n
}

func getStringCon() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userDB, passwordDB, hostDB, portDB, nameBaseDB)
}

// ListaNombres ????
func ListaNombres() {
	db, err := sql.Open("mysql", getStringCon())
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

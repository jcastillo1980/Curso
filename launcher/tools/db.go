package tools

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var (
	dbHost     string
	dbPort     int
	dbUser     string
	dbPassword string
	dbNameBase string
)

// init inicializa el modulo cargando los valores de las varibbles de conexión a la base por defecto
func init() {
	dbHost = "oficina.xuitec.com"
	dbPort = 3306
	dbUser = "golang"
	dbPassword = "golang"
	dbNameBase = "blx2"
}

// genCadenaConexion crea una cadena a partir de los parametros de conexion a ala base MySQL
func genCadenaConexion() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbNameBase)
}

//ProbarConexion  Hacemos una simple prueba de conexión con la base MySQL
func ProbarConexion() {
	fmt.Printf("Prueba Conexión Base MySQL [%s]\r\n", genCadenaConexion())

	db, err := sql.Open("mysql", genCadenaConexion())
	if err != nil {
		fmt.Println("ERROR OPEN: ", err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Println("ERROR PING:", err)
	} else {
		fmt.Println("BASE ABIERTA !!!")
	}

	valor := 0
	err = db.QueryRow("select count(*) from dev_con where id_dev in (1,2,3,6,7)", 1).Scan(&valor)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No ROws")
		} else {
			fmt.Println("ERROR QUERY:", err)
		}
	}
	fmt.Println("valor:", valor)
}

// deregisterTLSConfig ???
func deregisterTLSConfig(s string) {
	mysql.DeregisterTLSConfig(s)
}

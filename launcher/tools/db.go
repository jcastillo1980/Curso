package tools

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jcastillo1980/Curso/launcher/models"
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

	gp := models.DevGroup{}
	fmt.Println(gp)
}

// GetListaGroups retorna una lista de grupos de dispositivos
func GetListaGroups(isFix bool) ([]models.DevGroup, error) {

	strquery := ""
	if isFix == true {
		strquery = "select id,modomodem,medidas_inst,x_enable,x_task_ts,x_exe_pid,x_exe_ts,x_task_arg1,x_task_arg2,x_hm,x_nh,x_ti,x_tr,x_tm  from dev_group where medidas_inst='1' and x_enable='1' "
	} else {
		strquery = "select id,modomodem,medidas_inst,x_enable,x_task_ts,x_exe_pid,x_exe_ts,x_task_arg1,x_task_arg2,x_hm,x_nh,x_ti,x_tr,x_tm  from dev_group where medidas_inst='0' and x_enable='1' "
	}

	db, err := sql.Open("mysql", genCadenaConexion())
	if err != nil {
		log.Println("Error Open GetListaGroups:", err)
		return nil, err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Println("Error Ping GetListaGroups:", err)
		return nil, err
	}

	retornolista := make([]models.DevGroup, 0)
	row, err := db.Query(strquery)

	if err != nil {
		if err == sql.ErrNoRows {
			row.Close() // Puede que esto no sea correcto
			return retornolista, nil
		}

		log.Println("Error Query GetListaGroups:", err)
		return nil, err

	}

	defer row.Close()

	for row.Next() {
		elemento := models.DevGroup{}
		row.Scan(
			&elemento.Id,
			&elemento.Modomodem,
			&elemento.Medidas_inst,
			&elemento.X_enable,
			&elemento.X_task_ts,
			&elemento.X_exe_pid,
			&elemento.X_exe_ts,
			&elemento.X_task_arg1,
			&elemento.X_task_arg2,
			&elemento.X_hm,
			&elemento.X_nh,
			&elemento.X_ti,
			&elemento.X_tr,
			&elemento.X_tm,
		)

		retornolista = append(retornolista, elemento)

	}

	return retornolista, nil
}

// SetDbTaskRun Le pasas el id del grupo y indica task en marcha
func SetDbTaskRun(id int) {
	db, err := sql.Open("mysql", genCadenaConexion())
	if err != nil {
		log.Println("Error Open SetDbTaskRun:", err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Println("Error Ping SetDbTaskRun:", err)
		return
	}

	ts := time.Now().Unix()
	_, err = db.Exec("update dev_group set x_task_ts=? where id=? ", ts, id)
	if err != nil {
		log.Println("Error Ping SetDbTaskRun:", err)
	}
}

// SetDbTaskStop Le pasas el id del grupo y indica task parada. !!! tal vez habria que poner el exe tb parado???
func SetDbTaskStop(id int) {
	db, err := sql.Open("mysql", genCadenaConexion())
	if err != nil {
		log.Println("Error Open SetDbTaskStop:", err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Println("Error Ping SetDbTaskStop:", err)
		return
	}

	_, err = db.Exec("update dev_group set x_task_ts=0 where id=? ", id)
	if err != nil {
		log.Println("Error Ping SetDbTaskStop:", err)
	}
}

// SetDbExeRun Le pasas el id del grupo y indica que en una tarea se ejecuta un pid
func SetDbExeRun(id int, pid int) {
	db, err := sql.Open("mysql", genCadenaConexion())
	if err != nil {
		log.Println("Error Open SetDbExeRun:", err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Println("Error Ping SetDbExeRun:", err)
		return
	}

	ts := time.Now().Unix()

	_, err = db.Exec("update dev_group set x_exe_pid=?,x_exe_ts=?  where id=? ", pid, ts, id)
	if err != nil {
		log.Println("Error Ping SetDbExeRun:", err)
	}
}

// SetDbExeStop Le pasas el id del grupo y indica ejecutable parado
func SetDbExeStop(id int) {
	db, err := sql.Open("mysql", genCadenaConexion())
	if err != nil {
		log.Println("Error Open SetDbExeRun:", err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Println("Error Ping SetDbExeRun:", err)
		return
	}

	_, err = db.Exec("update dev_group set x_exe_pid=0,x_exe_ts=0  where id=? ", id)
	if err != nil {
		log.Println("Error Ping SetDbExeRun:", err)
	}
}

// deregisterTLSConfig ???
func deregisterTLSConfig(s string) {
	mysql.DeregisterTLSConfig(s)
}

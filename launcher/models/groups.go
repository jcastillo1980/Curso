package models

import (
	"fmt"
)

// DevGroup  ??????
type DevGroup struct {
	Id           int
	Modomodem    int
	Medidas_inst int
	X_enable     int
	X_task_ts    int64 
	X_exe_pid    int
	X_exe_ts     int64
	X_task_arg1  string
	X_task_arg2  string
	X_hm         int
	X_nh         int
	X_ti         int
	X_tr         int
	X_tm         int
}

func (d DevGroup) String() string {
	retorno := ""
	retorno = retorno + fmt.Sprintf("id:[%8d]---------------------------------------------------------------------\r\n", d.Id)
	retorno = retorno + fmt.Sprintf("modomodem:%d  medidas_inst:%d  x_enable:%d   x_task_ts:%d  x_exe_pid:%d\r\n", d.Modomodem, d.Medidas_inst, d.X_enable, d.X_task_ts, d.X_exe_pid)
	retorno = retorno + fmt.Sprintf("x_exe_ts:%d   x_hm:%d  x_nh:%d  x_ti:%d  x_tr:%d  x_tm:%d\r\n", d.X_exe_ts, d.X_hm, d.X_nh, d.X_ti, d.X_tr, d.X_tm)
	retorno = retorno + fmt.Sprintf("x_task_arg1:[%s]   x_task_arg1:[%s]\r\n", d.X_task_arg1, d.X_task_arg1)
	retorno = retorno + fmt.Sprintf("----------------------------------------------------------------------------------\r\n")

	return retorno
}

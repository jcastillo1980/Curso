package ElementExe

//ElementExe ??????
type ElementExe struct {
	NameExe  string
	NamePath string
	TR       int
	TI       int
	TM       int
	Arg1     string
	Arg2     string
	Arg3     string
	Finaliza chan int
}

//Ejecuta ejecuta una tarea de forma indefinida hasta que se lanza "Finaliza<-1"
func (e *ElementExe) Ejecuta() {

}

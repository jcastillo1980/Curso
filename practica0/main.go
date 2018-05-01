package main

import (
	"net"
	"os"
)

// 0 1 2 3 4 5
// 2 3 4 5
// 1 2 3 4 5
func remove(buff []int, indice int) []int {
	copy(buff[indice:], buff[indice+1:])
	return buff[:len(buff)-1]
}

func main() {
	/*fmt.Println("Esto es una prueba")
	pp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(pp)
	fmt.Println(remove(pp, 1))*/
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
}

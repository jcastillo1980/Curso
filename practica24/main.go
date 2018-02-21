package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	mapa := make(map[int]string)
	var keyss []int
	fmt.Println("Probando cosas")

	for _, v := range os.Environ() {
		index := strings.Index(v, "=")
		if index >= 0 {
			key := v[0:index]
			valor := v[index+1:]
			//fmt.Printf("[%s]-->[%s]\r\n", key, valor)
			index2 := strings.Index(key, "NOSTOP_ARG")
			if index2 >= 0 {
				var argn int
				fmt.Sscanf(key, "NOSTOP_ARG%d", &argn)
				mapa[argn] = valor
				keyss = append(keyss, argn)
			}
		}
	}

	sort.Ints(keyss)

	for _, vv := range keyss {
		fmt.Printf("[%d]------>[%s]\r\n", vv, mapa[vv])
	}
}

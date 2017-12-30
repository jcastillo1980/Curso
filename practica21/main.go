package main

import (
	"time"

	"github.com/kr/pretty"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	ct := time.After(10 * time.Second)
	go func() {
		gg := bson.M{
			"estado":    "algo",
			"valor":     33.3,
			"mas_cosas": true,
			"objecto": bson.M{
				"valor1": 44,
				"esto":   "mall mallll",
			},
		}
		pretty.Println(gg)
		gg["estado"] = "caca caca cac "
		pretty.Println(gg)
	}()
	<-ct
}

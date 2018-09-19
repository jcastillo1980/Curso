package main

import (
	"log"

	pt "github.com/kr/pretty"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Phone string
	//Timestamp time.Time
}

func main() {
	log.Println("Prueba conexion mongo")

	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Username: "root",
		Password: "cst2014C",
		Addrs:    []string{"valley.xuitec.com"},
	})
	if err != nil {
		log.Panicln(err)
	}
	defer s.Close()

	s.SetMode(mgo.Monotonic, true)

	c := s.DB("prueba").C("people")

	/*err = c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
		&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})

	if err != nil {
		log.Panicln(err)
	}*/

	var pp []Person
	err = c.Find(bson.M{"name": "Cla"}).All(&pp)
	if err != nil {
		log.Panicln(err)
	}

	pt.Printf("%# v\r\n", pp)

}

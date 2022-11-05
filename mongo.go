package main

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name  string
	Phone string
}

func mongoConnect() {
	session, err := mgo.Dial("localhost, server2.example.com")
	checkErr(err)
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 933 939 39394"},
		&Person{"Cla", "+33 339 3939394 39"})
	checkErr(err)

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	checkErr(err)

	fmt.Println("Name: ", result.Name, "Phone: ", result.Phone)
}

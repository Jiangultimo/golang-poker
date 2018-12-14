package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//Optional. Switch the session to a monotonic behavior
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("users")
	err = c.Insert(&Person{"Hing", "17682340306"}, &Person{"Neras", "18203096625"})
	if err != nil {
		log.Fatal(err)
	}
	 result := Person{}
	 err = c.Find(bson.M{"name": "Hing"}).One(&result)
	 if err != nil {
	 	log.Fatal(err)
	 }
	 fmt.Println("Phone", result.Phone)
}
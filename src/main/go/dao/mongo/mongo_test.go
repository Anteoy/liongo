package mongo

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
)

func TestMongo(t *testing.T) {
	c := Session.DB("test").C("people")
	err := c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
	defer Session.Close()
	cc := Session.DB("test2").C("people")
	err = cc.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result2 := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result2)
	if err != nil {
		log.Fatal(err)
	}

}

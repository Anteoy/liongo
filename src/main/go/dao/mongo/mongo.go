package mongo

import (
	"gopkg.in/mgo.v2"
)

type Person struct {
	Name  string
	Phone string
}

// 数据连接
var Session *mgo.Session

func init(){
	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	// mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
	var err error
	Session, err = mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)
}

func Mgo() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

}

package mongo

import (
	"gopkg.in/mgo.v2"
	"os"
)

type Person struct {
	Name  string
	Phone string
}

type DataStore struct {
	session *mgo.Session
}

// 数据连接
var Session *mgo.Session

func init() {
	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	// mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
	var err error
	if os.Getenv("liongo_env") == "online" {
		Session, err = mgo.Dial("mongodb://localhost:27017")
		if err != nil {
			panic(err)
		}
	} else if os.Getenv("liongo_env") == "compose" { //compose online
		Session, err = mgo.Dial("mongodb://mongodb:27017")
		if err != nil {
			panic(err)
		}
	} else { //compose local
		Session, err = mgo.Dial("mongodb://127.0.0.1:27017")
		if err != nil {
			panic(err)
		}
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

func GetMongoSession(ch chan *mgo.Session) {
	session := Session.Copy()
	ch <- session //error &session
}

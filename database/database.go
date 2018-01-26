package database

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
)

type Database struct {
	session  *mgo.Session
	database *mgo.Database
}

func (d *Database) Close() {
	d.session.Close()
}

func (d *Database) getCollection(name string) (coll *Collection) {
	coll = &Collection{
		*d.database.C(name),
		d,
	}
	return
}

func (d *Database) Hosts() (coll *Collection) {
	coll = d.getCollection("hosts")
	return
}

func Connect() (err error) {
	Session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		log.Fatalln(err)
		return
	}

	Session.SetMode(mgo.Strong, true)
	return
}

func GetDatabase() (db *Database) {
	session := Session.Copy()
	database := session.DB("pritunl")

	db = &Database{
		session:  session,
		database: database,
	}

	return
}

func init() {
	for {
		err := Connect()
		if err != nil {
			log.Fatalln(err)
		} else {
			break
		}

		time.Sleep(1 * time.Second)
	}
}

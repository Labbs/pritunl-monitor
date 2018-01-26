package hosts

import (
	"time"

	"github.com/Labbs/pritunl-monitor/database"
	"gopkg.in/mgo.v2/bson"
)

type Host struct {
	Id             string    `bson:"_id"`
	Name           string    `bson:"name"`
	Status         string    `bson:"status"`
	StartTimestamp time.Time `bson:"start_timestamp"`
	ThreadCount    int       `bson:"thread_count"`
	CpuUsage       float64   `bson:"cpu_usage"`
	MemUsage       float64   `bson:"mem_usage"`
	ServerCount    int       `bson:"server_count"`
	DeviceCount    int       `bson:"device_count"`
}

func GetHost(db *database.Database) (host *Host, err error) {
	coll := db.Hosts()
	host = &Host{}

	err = coll.Find(bson.M{"status": "online"}).One(host)
	if err != nil {
		err = database.ParseError(err)
		return
	}

	return
}

func GetHosts(db *database.Database) (hosts []*Host, err error) {
	coll := db.Hosts()
	hosts = []*Host{}

	cursor := coll.Find(nil).Iter()

	err = cursor.All(&hosts)
	if err != nil {
		err = database.ParseError(err)
		return
	}

	return
}

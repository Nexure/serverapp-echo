package cassandra

import (
  "github.com/gocql/gocql"
  log "github.com/sirupsen/logrus"
)

var Session *gocql.Session

func init() {

  Session = connect()

}

func connect()(*gocql.Session){
  var err error

  log.Info("Start Setup Cassandra!")
  cluster := gocql.NewCluster("cassandra")
  cluster.Keyspace = "system"
  //cluster.Consistency = gocql.Quorum
  session, err := cluster.CreateSession()
  if err != nil {
    log.Fatal(err)
    connect()
  }
  if session != nil {
    log.Info("Connection to Cassandra: Established")
  }
  return session
}

func Setup(){
  session := connect()
  defer session.Close()

  if err := session.Query(`
  CREATE KEYSPACE IF NOT EXISTS example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
  `).Exec(); err != nil {
    log.Fatal(err)
  }


  if err := session.Query(`
  CREATE TABLE IF NOT EXISTS example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));`).Exec(); err != nil {
    log.Fatal(err)
  }
  if err := session.Query(`CREATE INDEX IF NOT EXISTS ON example.tweet(timeline);`).Exec(); err != nil {
    log.Fatal(err)
  }

  // insert a tweet
  if err := session.Query(`INSERT INTO example.tweet (timeline, id, text) VALUES (?, ?, ?)`,
  "me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
    log.Fatal(err)
  }

}

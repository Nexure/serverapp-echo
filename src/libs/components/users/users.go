package users

import (
  "net/http"
  "github.com/labstack/echo"
  "github.com/rewiko/app/libs/cassandra"
  "github.com/gocql/gocql"
)


type User struct {
	ID        gocql.UUID `json:"id"`
	Text     string `json:"text"`
	TimeLine  string `json:timeline"`
}

func List(c echo.Context) error {

 var userList []User
  m := map[string]interface{}{}

  query := "SELECT * from example.tweet;"
  iterable := cassandra.Session.Query(query).Iter()
  for iterable.MapScan(m) {
    userList = append(userList, User{
      ID: m["id"].(gocql.UUID),
      Text: m["text"].(string),
      TimeLine: m["timeline"].(string),
    })
    m = map[string]interface{}{}
  }

  return c.JSON(http.StatusOK, userList)
}

package main

import (
  "net/http"
  "github.com/facebookgo/grace/gracehttp"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  //"github.com/rewiko/app/libs/admin"
  //"github.com/rewiko/app/libs/config"
  "github.com/rewiko/app/libs/log"
  "github.com/rewiko/app/libs/cassandra"
  "github.com/rewiko/app/libs/components/users"
)

type Person struct {
  Name  string
  Phone string
}

func main() {
  e := echo.New()

  //CORS
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
  }))

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  e.Use(middleware.RequestID())


  e.GET("/", func(c echo.Context) error {
    //fmt.Println("Phone:", result)

    return c.JSON(http.StatusOK, "testwf")
  })
  e.GET("/users", users.List)

  //admin.Main(r)
  //admin.Main(r, r.Group("/admin"))

  //fmt.Println("Routes: ", r.Routes())
  //config.SetConfig()
  log.Setup()
  cassandra.Setup()
  //jsonapi.Run()

  // Start server
  e.Server.Addr = ":8080"

  // Serve it like a boss
  e.Logger.Fatal(gracehttp.Serve(e.Server))
}


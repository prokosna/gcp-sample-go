// +build !appengine,!appenginevm

package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/prokosna/gcp-sample-go/lib"
)

func main() {
	e := lib.CreateMux()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("/", "public")
	e.Logger.Fatal(e.Start(":8080"))
}

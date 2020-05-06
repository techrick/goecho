package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	fmt.Printf("my app %s, commit %s, built at %s by %s", version, commit, date, builtBy)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World or not!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

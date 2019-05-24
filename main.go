package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/temp-go-dev/sample-api-crud/service"
	"github.com/temp-go-dev/sample-module"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, service.GetMessage()+sample.Sample())
	})
	e.Logger.Fatal(e.Start(":1323"))
}

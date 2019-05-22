package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/temp-go-dev/sample-api-crud/service"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, service.GetMessage())
	})
	e.Logger.Fatal(e.Start(":1323"))
}

// func GetMessage() string {
// 	return "Hellow, Porld!"
// }

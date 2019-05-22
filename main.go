package main

import (
	"net/http"
	
	"github.com/labstack/echo"

	sw "./service"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, sw.GetMessage())
	})
	e.Logger.Fatal(e.Start(":1323"))
}

// func getMessage() string {
// 	return "Hellow, Porld!"
// }

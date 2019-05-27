package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/temp-go-dev/sample-api-crud/model"
	"github.com/temp-go-dev/sample-api-crud/service"
	"github.com/temp-go-dev/sample-module"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, sample.Sample())
	})

	// users endpoint
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users", updateUser)
	e.DELETE("/users", deleteUser)

	// todos endpoint
	e.POST("/todos", createTodo)
	e.GET("/todos/:id", getTodo)
	e.PUT("/todos", updateTodo)
	e.DELETE("/todos", deleteTodo)

	e.Logger.Fatal(e.Start(":1323"))
}

func createUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, service.CreateUser(&user))
}
func getUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
func updateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
func deleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func createTodo(c echo.Context) error {
	return c.JSON(http.StatusCreated, "")
}
func getTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
func updateTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
func deleteTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

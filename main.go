package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/temp-go-dev/sample-api-crud/model"
	"github.com/temp-go-dev/sample-api-crud/service"
	"github.com/temp-go-dev/sample-module"
)

// main REST-APIサーバーの起動
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, sample.Sample())
	})

	// users endpoint
	e.POST("/users", createUser)
	e.GET("/users/:id", readUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// todos endpoint
	e.POST("/todos", createTodo)
	e.GET("/todos/:id", getTodo)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)

	e.Logger.Fatal(e.Start(":1323"))
}

func createUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	// execute business logic
	service.CreateUser(&user)
	return c.JSON(http.StatusCreated, user)
}
func readUser(c echo.Context) error {
	id := c.Param("id")
	user := service.ReadUser(id)
	return c.JSON(http.StatusOK, user)
}
func updateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	userId := c.Param("id")
	service.UpdateUser(userId, &user)
	return c.NoContent(http.StatusOK)
}
func deleteUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	userId := c.Param("id")
	service.DeleteUser(userId, &user)
	return c.NoContent(http.StatusOK)
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

package main

import (
	"net/http"

	controllers "echo/controllers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

func main() { //mendefinisikan routes yang ada di program kita
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", controllers.GetAllUser)
	e.POST("/users", controllers.AddUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	e.Logger.Fatal(e.Start(":8000"))

}

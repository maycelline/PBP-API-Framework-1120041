package main

import (
	"net/http"

	controllers "echo/controllers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() { //mendefinisikan routes yang ada di program kita
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", controllers.Authenticate(echo.HandlerFunc(controllers.GetAllUser), 0))
	e.POST("/users", controllers.Authenticate(echo.HandlerFunc(controllers.AddUser), 1))
	e.PUT("/users/:id", controllers.Authenticate(echo.HandlerFunc(controllers.UpdateUser), 1))
	e.DELETE("/users/:id", controllers.Authenticate(echo.HandlerFunc(controllers.DeleteUser), 2))

	e.POST("/login", controllers.Login)
	e.POST("/logout", controllers.Logout)
	e.Logger.Fatal(e.Start(":8000"))

}

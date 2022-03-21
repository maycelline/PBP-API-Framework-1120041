package main

import (
	"net/http"

	controllers "echo/controllers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() { //mendefinisikan routes yang ada di program kita
	router := echo.New()
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/users", controllers.Authenticate(echo.HandlerFunc(controllers.GetAllUser), 0))
	router.POST("/users", controllers.Authenticate(echo.HandlerFunc(controllers.AddUser), 1))
	router.PUT("/users/:id", controllers.Authenticate(echo.HandlerFunc(controllers.UpdateUser), 1))
	router.DELETE("/users/:id", controllers.Authenticate(echo.HandlerFunc(controllers.DeleteUser), 2))

	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)
	router.Logger.Fatal(router.Start(":8000"))

}

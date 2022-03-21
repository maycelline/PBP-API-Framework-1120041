package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	db := Connect()
	defer db.Close()

	username := c.FormValue("username")
	password := c.FormValue("password")

	query := "SELECT id, username, address, hobby, type FROM users WHERE username = '" + username + "' AND password = '" + password + "'"
	row := db.QueryRow(query)

	var user User

	if err := row.Scan(&user.Id, &user.Username, &user.Address, &user.Hobby, &user.Type); err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, "Login failed. Please try again")
	} else {
		generateToken(c, user.Id, user.Username, user.Type)
		return c.JSON(http.StatusOK, "Login success!")
	}
}

func GetAllUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	var user User
	var users []User
	var response UsersResponse

	sqlStatement := "SELECT id, username, address, hobby, type FROM users"

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error Query")
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Address, &user.Hobby, &user.Type)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error Data")
		}

		users = append(users, user)
	}

	if len(users) <= 0 {
		return c.JSON(http.StatusNotFound, "No Data Available")
	} else {
		response.Message = "success"
		response.Data = users
		return c.JSON(http.StatusOK, response)
	}
}

func Logout(c echo.Context) error {
	ResetUserToken(c)
	return c.JSON(http.StatusOK, "Successfully Logout!")
}

func AddUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	username := c.FormValue("username")
	address := c.FormValue("address")
	hobby := c.FormValue("hobby")
	password := c.FormValue("password")

	if len(username) == 0 || len(address) == 0 || len(hobby) == 0 || len(password) == 0 {
		return c.JSON(http.StatusBadRequest, "Please input all fields!")
	}

	_, err := db.Exec("INSERT INTO users(username, address, hobby, password, type) VALUES(?,?,?,?,0)",
		username,
		address,
		hobby,
		password,
	)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Can not inser new user")
	} else {
		return c.JSON(http.StatusCreated, "New user inserted!")
	}
}

func UpdateUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	userid := c.Param("id")
	username := c.FormValue("username")
	address := c.FormValue("address")
	hobby := c.FormValue("hobby")

	if len(username) == 0 || len(address) == 0 || len(hobby) == 0 {
		return c.JSON(http.StatusBadRequest, "Please input all fields")
	}

	result, err := db.Exec("UPDATE users SET username = ?, address = ?, hobby = ? WHERE id = ?",
		username,
		address,
		hobby,
		userid,
	)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Can not update user")
	} else {
		number, _ := result.RowsAffected()
		if number == 0 {
			return c.JSON(http.StatusNotFound, "User with id "+userid+" not found or you made no changes at all")
		}
		return c.JSON(http.StatusOK, "User Updated!")
	}
}

func DeleteUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	userid := c.Param("id")

	result, err := db.Exec("DELETE FROM users WHERE id = ?",
		userid,
	)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Not success")
	} else {
		number, _ := result.RowsAffected()
		if number == 0 {
			return c.JSON(http.StatusNotFound, "User with id "+userid+" not found")
		}
		return c.JSON(http.StatusOK, "User Deleted!")
	}
}

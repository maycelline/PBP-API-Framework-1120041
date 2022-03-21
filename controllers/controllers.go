package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {

	var user User
	var users []User
	var userResponse UsersResponse
	var response Response

	db := Connect()
	defer db.Close()

	sqlStatement := "SELECT * FROM users"

	rows, err := db.Query(sqlStatement)

	if err != nil {
		response.Message = "Error with Query"
		return c.JSON(http.StatusInternalServerError, response)
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Address, &user.Hobby)

		if err != nil {
			response.Message = "Error with Data"
			return c.JSON(http.StatusInternalServerError, response)
		}

		users = append(users, user)
	}

	userResponse.Message = "Success"
	userResponse.Data = users

	return c.JSON(http.StatusOK, userResponse)
}

func AddUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	name := c.FormValue("name")
	address := c.FormValue("address")
	hobby := c.FormValue("hobby")
	var response Response

	if len(name) == 0 || len(address) == 0 || len(hobby) == 0 {
		response.Message = "Please Input All Fields"
		return c.JSON(http.StatusBadRequest, response)
	}

	_, err := db.Exec("INSERT INTO users(name, address, hobby) VALUES(?,?,?)",
		name,
		address,
		hobby,
	)

	if err != nil {
		log.Println()
		response.Message = "Not success"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Message = "New user inserted!"
		return c.JSON(http.StatusCreated, response)
	}
}

func UpdateUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	userid := c.Param("id")
	name := c.FormValue("name")
	address := c.FormValue("address")
	hobby := c.FormValue("hobby")
	var response Response

	if len(name) == 0 || len(address) == 0 || len(hobby) == 0 {
		response.Message = "Please Input All Fields"
		return c.JSON(http.StatusBadRequest, response)
	}

	result, err := db.Exec("UPDATE users SET name = ?, address = ?, hobby = ? WHERE id = ?",
		name,
		address,
		hobby,
		userid,
	)

	if err != nil {
		log.Println(err)
		response.Message = "Not success"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		number, _ := result.RowsAffected()
		if number == 0 {
			response.Message = "User with id " + userid + " not found or you made no changes at all"
			return c.JSON(http.StatusOK, response)
		}
		response.Message = "User Updated!"
		return c.JSON(http.StatusOK, response)
	}
}

func DeleteUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	userid := c.Param("id")
	var response Response

	result, err := db.Exec("DELETE FROM users WHERE id = ?",
		userid,
	)

	if err != nil {
		log.Println(err)
		response.Message = "Not success"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		number, _ := result.RowsAffected()
		if number == 0 {
			response.Message = "User with id " + userid + " not found"
			return c.JSON(http.StatusOK, response)
		}
		response.Message = "User Deleted!"
		return c.JSON(http.StatusOK, response)
	}
}

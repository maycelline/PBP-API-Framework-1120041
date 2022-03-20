package controllers

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Hobby   string `json:"hobby"`
}

type UsersResponse struct {
	Message string `json:"id"`
	Data    []User `json:"users"`
}

type Response struct {
	Message string `json:"message"`
}

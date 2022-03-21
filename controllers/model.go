package controllers

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Hobby    string `json:"hobby"`
	Type     int    `json:"usertype"`
}

type UsersResponse struct {
	Message string `json:"message"`
	Data    []User `json:"users"`
}

package models

type User struct {
	Id       int `json:"id" orm:"column(id);pk"`
	Username string
	Password string
}

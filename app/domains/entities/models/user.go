package models

import "github.com/vality-prototype/vality-utility-service/configs"

type User struct {
	configs.SQLModel `json:",inline"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	FullName         string `json:"full_name"`
}

func (User) TableName() string {
	return "users"
}

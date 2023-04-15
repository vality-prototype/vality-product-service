package models

import "github.com/vality-prototype/vality-utility-service/configs"

type Product struct {
	configs.SQLModel `json:",inline"`
	Name             string `json:"name"`
	Price            int    `json:"price"`
}

package main

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id       int
	Name     string
	Password string
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
}

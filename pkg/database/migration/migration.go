package main

import (
	"github.com/mrmertkose/gotth-auth/model"
	"github.com/mrmertkose/gotth-auth/pkg/database"
	"github.com/mrmertkose/gotth-auth/pkg/password"
)

func main() {
	database.Connect()
	db := database.Connection()
	db.AutoMigrate(&model.User{})

	pass, _ := password.HashPassword("123456")
	user := &model.User{
		Name:     "John Doe",
		Email:    "mail@mail.com",
		Password: pass,
	}

	db.Create(&user)
}

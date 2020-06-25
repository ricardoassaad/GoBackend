package controllers

import (
	"github.com/ricardoassaad/GoBackend/json"
	"github.com/ricardoassaad/GoBackend/models"
)

func UsersIndex() []byte {
	users := models.All()

	view := json.ListUsers(users)

	return view
}

func UsersShow(id string) []byte {
	user := models.ShowUser(id)

	view := json.ShowUser(user)

	return view
}

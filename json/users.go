package json

import (
	"bytes"
	"encoding/json"
	"github.com/ricardoassaad/GoBackend/models"
)

func ListUsers(users []models.User) []byte {
	response := new(bytes.Buffer)

	json.NewEncoder(response).Encode(users)

	return response.Bytes()
}

func ShowUser(user models.User) []byte {
	response := new(bytes.Buffer)

	json.NewEncoder(response).Encode(user)

	return response.Bytes()
}

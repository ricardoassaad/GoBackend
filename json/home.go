package json

import (
	"bytes"
	"encoding/json"
)

func Index() []byte {
	response := new(bytes.Buffer)

	json.NewEncoder(response).Encode("Welcome")

	return response.Bytes()
}

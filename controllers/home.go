package controllers

import (
	"github.com/ricardoassaad/GoBackend/json"
)

func Index() []byte {
	view := json.Index()

	return view
}

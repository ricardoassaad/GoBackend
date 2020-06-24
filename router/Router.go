package router

import (
	Users "github.com/ricardoassaad/GoBackend/controllers/UsersController"
	"net/http"
)

func Route() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(Users.Index()))
	})
	http.ListenAndServe(":8000", nil)
}

package router

import (
	"github.com/gorilla/mux"
	"github.com/ricardoassaad/GoBackend/controllers"
	"net/http"
	"os"
)

func Route() {
	router := mux.NewRouter()
	router.Host(os.Getenv("APP_HOST"))

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write(controllers.UsersIndex())
	}).Methods("GET")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write(controllers.UsersShow(vars["id"]))
	}).Methods("GET")

	http.ListenAndServe(":"+os.Getenv("APP_PORT"), router)
}

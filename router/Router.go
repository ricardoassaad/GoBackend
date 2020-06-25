package router

import (
	"github.com/gorilla/mux"
	"github.com/ricardoassaad/GoBackend/controllers"
	"github.com/ricardoassaad/GoBackend/middlewares"
	"net/http"
	"os"
)

func Route() {
	router := mux.NewRouter()
	router.Host(os.Getenv("APP_HOST"))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(controllers.Index())
	}).Methods("GET")
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write(controllers.UsersIndex())
	}).Methods("GET")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write(controllers.UsersShow(vars["id"]))
	}).Methods("GET")

	http.ListenAndServe(":"+os.Getenv("APP_PORT"), middlewares.RemoveTrailingSlash(router))
}

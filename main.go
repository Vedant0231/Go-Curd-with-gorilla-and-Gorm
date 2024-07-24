package main

import (
	src "go_curd/src"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func router() {
	r := mux.NewRouter()

	r.HandleFunc("/get/user", src.Getusers).Methods("GET")
	r.HandleFunc("/get/user/{id}", src.Getuser).Methods("GET")
	r.HandleFunc("/create/user", src.CreateUser).Methods("POST")
	r.HandleFunc("/update/users/{id}", src.Updateuser).Methods("UPDATE")
	r.HandleFunc("/delete/users/{id}", src.Deleteuser).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:3000", r))
}

func main() {

	src.Initalmigration()
	router()

}

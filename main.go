package main

import (
	"assigment-2-scalable-go/config"
	"assigment-2-scalable-go/service"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func start(db *gorm.DB) {
	r := mux.NewRouter()
	fmt.Println("starting server...")
	// open service for order
	serviceOrder := service.Order{DB: db}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "J")
	}).Methods("GET")

	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Ja")
	}).Methods("GET")

	r.HandleFunc("/api/create", func(w http.ResponseWriter, r *http.Request) {
		s := serviceOrder.InsertOrder(w, r)
		service.SendResponse(w, s)
	}).Methods("POST")


	http.Handle("/", r)
}
func main() {

	db := config.DbInit()
	start(db)

	err := http.ListenAndServe(":8080", nil)

	fmt.Print("server running on ")
	if err != nil {
		panic(err.Error())
	}
}

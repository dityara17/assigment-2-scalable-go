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
		_, _ = fmt.Fprint(w, "Hello")
	}).Methods("GET")

	r.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		s := serviceOrder.InsertOrder(w, r)
		service.SendResponse(w, s)
	}).Methods("POST")

	r.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		s := serviceOrder.UpdateOrder(w, r)
		service.SendResponse(w, s)
	}).Methods("PUT")

	r.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		s := serviceOrder.GetOrders(w, r)
		service.SendResponse(w, s)
	}).Methods("GET")

	r.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		s := serviceOrder.GetOrder(w, r)
		service.SendResponse(w, s)
	}).Methods("GET")

	r.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		s := serviceOrder.DeleteOrder(w, r)
		service.SendResponse(w, s)
	}).Methods("DELETE")

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

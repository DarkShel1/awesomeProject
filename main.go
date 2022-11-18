package main

import (
	"awesomeProject/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "croftsky1",
			DB:         "golang_api",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	log.Println("HTTP server starts on port 8090")
	router := mux.NewRouter()

	router.HandleFunc("/create", createBalance).Methods("POST")

	router.HandleFunc("/get/{id}", getBalanceByID).Methods("GET")

	router.HandleFunc("/update/{id}", updateBalanceByID).Methods("PUT")

	router.HandleFunc("/delete/{id}", deleteBalanceByID).Methods("DELETE")

	router.HandleFunc("/create_reservation", createReservation).Methods("POST")

	router.HandleFunc("/get_reservation/{id}", getReservationByID).Methods("GET")

	router.HandleFunc("/update_reservation/{id}", updateReservationByID).Methods("PUT")

	router.HandleFunc("/delete_reservation/{id}", deleteReservationByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", router))
}

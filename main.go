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
			Password:   "root",
			DB:         "dbName",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	log.Println("HTTP server starts on port 8090")
	router := mux.NewRouter()

	router.HandleFunc("/create", createBalance).Methods("POST") //создание баланса юзеров

	router.HandleFunc("/get/{id}", getBalanceByID).Methods("GET") //получение баланса юзеров по ID

	router.HandleFunc("/update/{id}", updateBalanceByID).Methods("PUT") //изменение баланса юзеров по ID

	router.HandleFunc("/delete/{id}", deleteBalanceByID).Methods("DELETE") //удаление баланса юзеров по ID

	router.HandleFunc("/create_reservation", createReservation).Methods("POST") //создание резервирования суммы

	router.HandleFunc("/get_reservation/{id}", getReservationByID).Methods("GET") //получение резервирования суммы по ID

	router.HandleFunc("/update_reservation/{id}", updateReservationByID).Methods("PUT") //изменение резервирования суммы по ID

	router.HandleFunc("/delete_reservation/{id}", deleteReservationByID).Methods("DELETE") //удаление резервирования суммы по ID

	log.Fatal(http.ListenAndServe(":8090", router))
}

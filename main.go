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
			DB:         "DbName",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	log.Println("HTTP server starts on port 8090")
	router := mux.NewRouter()

	/*создание баланса юзеров*/
	router.HandleFunc("/create", createBalance).Methods("POST")

	/*получение баланса юзеров по ID*/
	router.HandleFunc("/get/{id}", getBalanceByID).Methods("GET")

	/*изменение баланса юзеров по ID*/
	router.HandleFunc("/update/{id}", updateBalanceByID).Methods("PUT")

	/*удаление баланса юзеров по ID*/
	router.HandleFunc("/delete/{id}", deleteBalanceByID).Methods("DELETE")

	/*создание резервирования суммы*/
	router.HandleFunc("/create_reservation", createReservation).Methods("POST")

	/*получение резервирования суммы по ID*/
	router.HandleFunc("/get_reservation/{id}", getReservationByID).Methods("GET")

	/*изменение резервирования суммы по ID*/
	router.HandleFunc("/update_reservation/{id}", updateReservationByID).Methods("PUT")

	/*удаление резервирования суммы по ID*/
	router.HandleFunc("/delete_reservation/{id}", deleteReservationByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", router))
}

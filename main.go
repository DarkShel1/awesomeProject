package main

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func createPerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	var person model.User
	err := json.Unmarshal(requestBody, &person)
	if err != nil {
		return
	}

	database.Connector.Create(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	er := json.NewEncoder(w).Encode(person)
	if er != nil {
		return
	}
}

func getPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person model.User
	database.Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		return
	}
}

func updatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	var person model.User
	err := json.Unmarshal(requestBody, &person)
	if err != nil {
		return
	}
	database.Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	er := json.NewEncoder(w).Encode(person)
	if er != nil {
		return
	}
}

func deletePersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person model.User
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "croftsky1",
			DB:         "mydbtest",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	log.Println("HTTP server starts on port 8090")
	router := mux.NewRouter()

	router.HandleFunc("/create", createPerson).Methods("POST")

	router.HandleFunc("/get/{id}", getPersonByID).Methods("GET")

	router.HandleFunc("/update/{id}", updatePersonByID).Methods("PUT")

	router.HandleFunc("/delete/{id}", deletePersonByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", router))
}

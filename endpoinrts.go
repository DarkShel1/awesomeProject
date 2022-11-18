package main

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func createBalance(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	var balance model.UserBalance
	err := json.Unmarshal(requestBody, &balance)
	if err != nil {
		return
	}

	database.Connector.Create(balance)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	er := json.NewEncoder(w).Encode(balance)
	if er != nil {
		return
	}
}

func getBalanceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var balance model.UserBalance
	database.Connector.First(&balance, key)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(balance)
	if err != nil {
		return
	}
}

func updateBalanceByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	var balance model.UserBalance
	err := json.Unmarshal(requestBody, &balance)
	if err != nil {
		return
	}
	database.Connector.Save(&balance)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	er := json.NewEncoder(w).Encode(balance)
	if er != nil {
		return
	}
}

func deleteBalanceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var balance model.UserBalance
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&balance)
	w.WriteHeader(http.StatusNoContent)
}

func createReservation(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	var reservation model.Reservation
	err := json.Unmarshal(requestBody, &reservation)
	if err != nil {
		return
	}

	database.Connector.Create(reservation)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	er := json.NewEncoder(w).Encode(reservation)
	if er != nil {
		return
	}
}

func getReservationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var balance model.UserBalance
	database.Connector.First(&balance, key)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(balance)
	if err != nil {
		return
	}
}

func updateReservationByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	var balance model.UserBalance
	err := json.Unmarshal(requestBody, &balance)
	if err != nil {
		return
	}
	database.Connector.Save(&balance)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	er := json.NewEncoder(w).Encode(balance)
	if er != nil {
		return
	}
}

func deleteReservationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var balance model.UserBalance
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&balance)
	w.WriteHeader(http.StatusNoContent)
}

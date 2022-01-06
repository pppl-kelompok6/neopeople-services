package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"neopeople-service/database"
	"neopeople-service/model"
	"net/http"

	"github.com/gorilla/mux"
)

func GetEventAll(w http.ResponseWriter, r *http.Request) {
	var event []model.Event

	err := database.Connector.Find(&event).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func GetEventByID(w http.ResponseWriter, r *http.Request) {
	var event []model.Event
	id := mux.Vars(r)
	key := id["id"]

	err := database.Connector.Preload("Attendance").Preload("EventOrder").Find(&event, "events.id = ?", key).Error
	if err != nil {
		fmt.Println("line 39")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTION")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var event model.Event

	err := json.Unmarshal(reqBody, &event)

	if err != nil {
		fmt.Println("line 60")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = database.Connector.Create(&event).Error
	if err != nil {
		fmt.Println("(line 69")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Event has been created")

}

func UpdateEventById(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var eventUpdate model.Event
	var event model.Event
	id := mux.Vars(r)
	key := id["id"]

	err := json.Unmarshal(reqBody, &eventUpdate)

	if err != nil {
		fmt.Println("line 92")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = database.Connector.First(&event, key).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = database.Connector.Model(&event).Updates(&eventUpdate).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Event has been updated")

}

func DeleteEventById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var event model.Event
	err := database.Connector.First(&event, id).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	database.Connector.Delete(&event)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Event has been deleted")
}

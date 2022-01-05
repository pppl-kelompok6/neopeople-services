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

func GetTeamAll(w http.ResponseWriter, r *http.Request) {
	var team []model.Team

	err := database.Connector.Find(&team).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}

func GetTeamByID(w http.ResponseWriter, r *http.Request) {
	var team []model.Team

	err := database.Connector.Preload("Attendance").Preload("Pantient").Find(&team).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var team model.Team

	json.Unmarshal(reqBody, &team)

	err := json.NewDecoder(r.Body).Decode(&team)

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = database.Connector.Create(team).Error
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Team has been created")

}

func UpdateTeamById(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var teamUpdate model.Team
	var team model.Team
	id := mux.Vars(r)
	key := id["id"]

	json.Unmarshal(reqBody, &teamUpdate)
	err := json.NewDecoder(r.Body).Decode(&teamUpdate)

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = database.Connector.First(&team, key).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = database.Connector.Model(&team).Updates(&teamUpdate).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Team has been updated")

}

func DeleteTeamById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var team model.Team
	err := database.Connector.First(&team, id).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	database.Connector.Delete(&team)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Team has been deleted")
}

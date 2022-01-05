package controllers

import (
	"encoding/json"
	"fmt"
	"neopeople-service/database"
	"neopeople-service/model"
	"net/http"
)

func GetFaqAll(w http.ResponseWriter, r *http.Request) {
	var faq []model.Faq

	err := database.Connector.Find(&faq).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(faq)
}

func GetFaqByID(w http.ResponseWriter, r *http.Request) {
	var faq []model.Faq

	err := database.Connector.Preload("Attendance").Preload("Pantient").Find(&faq).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(faq)
}

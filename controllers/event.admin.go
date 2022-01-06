package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"neopeople-service/cdn"
	"neopeople-service/database"
	"neopeople-service/model"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/api/uploader"
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

	file, _, err := r.FormFile("file")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("line 60")
		fmt.Println(err)
		return
	}

	event_name := r.FormValue("event_name")
	date := r.FormValue("date")
	started_at := r.FormValue("started_at")
	finish_at := r.FormValue("finish_at")
	price := r.FormValue("price")
	speaker := r.FormValue("speaker")
	speaker_job := r.FormValue("speaker_job")
	speaker_company := r.FormValue("speaker_company")
	description := r.FormValue("description")

	defer file.Close()

	cld, err := cdn.CdnSetting()
	var ctx = context.Background()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("line 82")
		fmt.Println("error open file", err)
		return
	}

	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("line 91")
		fmt.Println("error read file", err)
		return
	}

	cover := uploadResult.SecureURL
	fmt.Println(cover)
	var event model.Event

	event.EventName = event_name
	event.Cover = cover
	event.Date = date
	event.StartedAt = started_at
	event.Description = description
	event.SpeakerJob = speaker_job
	event.SpeakerCompany = speaker_company
	event.Speaker = speaker
	event.FinishAt = finish_at
	event.Price, _ = strconv.Atoi(price)

	// err = json.Unmarshal(reqBody, &event)
	fmt.Println("ini")
	fmt.Println(event)
	// if err != nil {
	// 	fmt.Println("line 60")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	// err = database.Connector.Create(&event).Error
	// if err != nil {
	// 	fmt.Println("(line 69")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }
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

package main

import (
	"fmt"
	"log"
	"neopeople-service/controllers"
	"neopeople-service/database"
	"net/http"

	"github.com/gorilla/mux"
)

func Services(router *mux.Router) {
	staticDir := "/static/"
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	router.HandleFunc("/events", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/events", controllers.GetEventAll).Methods("GET")
	router.HandleFunc("/events/{id}", controllers.GetEventByID).Methods("GET")
	router.HandleFunc("/events/{id}", controllers.UpdateEventById).Methods("PUT")
	router.HandleFunc("/events/{id}", controllers.DeleteEventById).Methods("DELETE")
}

func RouterStart() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println(`Running on port 8090`)
	Services(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func InitDB() {
	config :=
		database.Config{
			// ServerName: "34.69.20.223:3306",
			// User:       "staggingUser",
			// Pass:       "@WTx3GV^@7aJk9m2",
			// DB:         "neo_stagging",
			ServerName: "localhost:3306",
			User:       "root",
			Pass:       "",
			DB:         "neo_stagging",
		}
	connectionString := database.GetConnectionString(config)
	err := database.Connect((connectionString))
	if err != nil {
		panic(err.Error())
	}
	// database.Migrate()
}

func main() {
	InitDB()
	RouterStart()
}

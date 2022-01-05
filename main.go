package main

import (
	"fmt"
	"log"
	"neopeople-service/controllers"
	"neopeople-service/database"
	"neopeople-service/middleware"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Services(router *mux.Router) {
	staticDir := "/static/"
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	// auth
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	// events
	router.HandleFunc("/events", middleware.Authorization(controllers.CreateEvent)).Methods("POST")
	router.HandleFunc("/events", controllers.GetEventAll).Methods("GET")
	router.HandleFunc("/events/{id}", middleware.Authorization(controllers.GetEventByID)).Methods("GET")
	router.HandleFunc("/events/{id}", middleware.Authorization(controllers.UpdateEventById)).Methods("PUT")
	router.HandleFunc("/events/{id}", middleware.Authorization(controllers.DeleteEventById)).Methods("DELETE")

	// session
	router.HandleFunc("/session", middleware.Authorization(controllers.CreateSession)).Methods("POST")
	router.HandleFunc("/session", controllers.GetSessionAll).Methods("GET")
	router.HandleFunc("/session/{id}", middleware.Authorization(controllers.GetSessionByID)).Methods("GET")
	router.HandleFunc("/session/{id}", middleware.Authorization(controllers.UpdateSesionById)).Methods("PUT")
	router.HandleFunc("/session/{id}", middleware.Authorization(controllers.DeleteSessionById)).Methods("DELETE")

}

func RouterStart() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println(`Running on port 8090`)
	Services(router)
	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8090", loggedRouter))
}

func InitDB() {
	config :=
		database.Config{
			// 	ServerName: "34.69.20.223:3306",
			// 	User:       "staggingUser",
			// 	Pass:       "@WTx3GV^@7aJk9m2",
			// 	DB:         "neo_stagging",
			ServerName: "localhost:3300",
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

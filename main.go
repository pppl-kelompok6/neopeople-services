package main

import (
	"fmt"
	"log"
	"neopeople-service/database"
	"net/http"

	"github.com/gorilla/mux"
)

func Services(router *mux.Router) {
	staticDir := "/static/"
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

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
			ServerName: "34.69.20.223:3306",
			User:       "staggingUser",
			Pass:       "@WTx3GV^@7aJk9m2",
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

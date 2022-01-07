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
	"github.com/rs/cors"
)

func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT")
			w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, token")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func Services(router *mux.Router) {
	staticDir := "/static/"
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	// auth
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	// events
	router.HandleFunc("/events", middleware.Authorization(controllers.CreateEvent)).Methods("POST")
	router.HandleFunc("/events", controllers.GetEventAll).Methods("GET")
	router.HandleFunc("/events/{id}", controllers.GetEventByID).Methods("GET")
	router.HandleFunc("/events/{id}", middleware.Authorization(controllers.UpdateEventById)).Methods("PUT")
	router.HandleFunc("/events/{id}", middleware.Authorization(controllers.DeleteEventById)).Methods("DELETE")

	// session
	router.HandleFunc("/session", controllers.CreateSession).Methods("POST")
	router.HandleFunc("/session", controllers.GetSessionAll).Methods("GET")
	router.HandleFunc("/session/{id}", controllers.GetSessionByID).Methods("GET")
	router.HandleFunc("/session/{id}", middleware.Authorization(controllers.UpdateSesionById)).Methods("PUT")
	router.HandleFunc("/session/{id}", middleware.Authorization(controllers.DeleteSessionById)).Methods("DELETE")

	// pantient
	router.HandleFunc("/patient", controllers.GetPatientAll).Methods("GET")
	router.HandleFunc("/patient", controllers.CreatePatient).Methods("POST")
	router.HandleFunc("/patient/{id}", controllers.GetPatientByID).Methods("GET")
	router.HandleFunc("/patient/{id}", middleware.Authorization(controllers.UpdatePatientById)).Methods("PUT")
	router.HandleFunc("/patient/{id}", middleware.Authorization(controllers.DeletePatientById)).Methods("DELETE")

	// counselor
	router.HandleFunc("/counselor", controllers.GetCounselorAll).Methods("GET")
	router.HandleFunc("/counselor", controllers.CreateCounselor).Methods("POST")
	router.HandleFunc("/counselor/{id}", controllers.GetCounselorByID).Methods("GET")
	router.HandleFunc("/counselor/{id}", middleware.Authorization(controllers.UpdateCounselorById)).Methods("PUT")
	router.HandleFunc("/counselor/{id}", middleware.Authorization(controllers.DeleteCounselorById)).Methods("DELETE")

	// counselor
	router.HandleFunc("/attendance", controllers.GetAttendanceAll).Methods("GET")
	router.HandleFunc("/attendance", controllers.CreateAttendance).Methods("POST")
	router.HandleFunc("/attendance/{id}", controllers.GetAttendanceByID).Methods("GET")
	router.HandleFunc("/attendance/{id}", middleware.Authorization(controllers.UpdateAttendanceById)).Methods("PUT")
	router.HandleFunc("/attendance/{id}", middleware.Authorization(controllers.DeleteAttendanceById)).Methods("DELETE")

}

func RouterStart() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println(`Running on port 3001`)
	Services(router)
	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
	fmt.Println(loggedRouter)
	handle := cors.Default().Handler(router)
	fmt.Println(handle)
	log.Fatal(http.ListenAndServe(":3001", CORS(loggedRouter)))
}

func InitDB() {
	config :=
		database.Config{
			// 	ServerName: "34.69.20.223:3306",
			// 	User:       "staggingUser",
			// 	Pass:       "@WTx3GV^@7aJk9m2",
			// 	DB:         "neo_stagging",
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

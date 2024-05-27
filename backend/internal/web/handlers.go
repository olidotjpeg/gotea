package web

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var Database *sql.DB

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/teas", getTeas).Methods(http.MethodGet)
	myRouter.HandleFunc("/teas/status", getTeaStatus).Methods(http.MethodGet)
	// NOTE: Ordering is important here! This has to be defined before
	// the other `/tea` endpoint.
	myRouter.HandleFunc("/tea", createNewTea).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc("/tea/{id}", deleteTea).Methods(http.MethodDelete, http.MethodOptions)
	myRouter.HandleFunc("/tea/{id}", updateTea).Methods(http.MethodPut, http.MethodOptions)
	myRouter.HandleFunc("/tea/{id}", returnSingleTea).Methods(http.MethodGet)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost", "http://localhost:5173"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	srv := &http.Server{
		Handler: c.Handler(myRouter),
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

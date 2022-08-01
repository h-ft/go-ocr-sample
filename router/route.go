package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/handler"
)

// A function to handle the request and pass them to its corresponding router
func HandleRequests(handler *handler.Handler) {

	// Creates a new instance of a mux router
	muxRouter := mux.NewRouter().StrictSlash(true)
	// Replace http.HandleFunc with myRouter.HandleFunc
	muxRouter.HandleFunc("/", handler.HomePage)
	muxRouter.HandleFunc("/read", handler.ReadImage).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", muxRouter))
}

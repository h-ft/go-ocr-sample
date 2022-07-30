package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/otiai10/gosseract"
	"main.go/handler"
)

// A function to handle the request and pass them to its corresponding router
func HandleRequests(gosseractClient *gosseract.Client) {

	//api  handlers
	handler := handler.Handler{
		Gosseract: gosseractClient,
	}

	// Creates a new instance of a mux router
	muxRouter := mux.NewRouter().StrictSlash(true)
	// Replace http.HandleFunc with myRouter.HandleFunc
	muxRouter.HandleFunc("/", homePage)
	muxRouter.HandleFunc("/read", handler.ReadImage).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", muxRouter))
}

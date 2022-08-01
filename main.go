package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
	"main.go/functions"
	"main.go/handler"
	http "main.go/router"
)

func main() {
	go fmt.Println("Server started")

	handler := initialize()

	http.HandleRequests(handler)
}

func initialize() *handler.Handler {

	gosseract := gosseract.NewClient()
	functions := functions.New(gosseract)

	handler := handler.New(functions)

	return handler
}

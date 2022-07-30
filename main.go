package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
	http "main.go/router"
)

func main() {
	go fmt.Println("Server started")

	gosseract := gosseract.NewClient()

	http.HandleRequests(gosseract)
}

package http

import (
	"fmt"
	"net/http"
)

// Homepage default function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok!")
	fmt.Println("Endpoint: homePage")
}

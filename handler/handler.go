package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type (
	Handler struct {
		Functions Functions
	}

	Functions interface {
		ReadImage(ctx context.Context) error
	}
)

func New(functions Functions) *Handler {
	return &Handler{
		Functions: functions,
	}
}

func (h *Handler) ReadImage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	err := h.Functions.ReadImage(ctx)
	if err != nil {
		logrus.Errorf("Error reading image: %v", err)
		handleError(w, r)
		return
	}
	// add return type
	return
}

func (h *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok!")
	fmt.Println("Endpoint: homePage")
}

package handler

import (
	"net/http"

	"github.com/otiai10/gosseract"
)

type Handler struct {
	Gosseract *gosseract.Client
}

func (h *Handler) ReadImage(w http.ResponseWriter, r *http.Request) {
	return
}

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func handleError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Internal Error"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		logrus.Errorf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	w.Write(jsonResp)
	return
}

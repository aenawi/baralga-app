package util

import (
	"encoding/json"
	"log"
	"net/http"

	"schneider.vip/problem"
)

func RenderJSON(w http.ResponseWriter, jsonModel interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonModel)
	if err != nil {
		http.Error(w, problem.New(problem.Wrap(err)).JSONString(), http.StatusInternalServerError)
	}
}

func RenderProblemJSON(w http.ResponseWriter, isProduction bool, err error) {
	log.Printf("internal server error: %s", err)

	if !isProduction {
		http.Error(w, problem.New(problem.Title("internal server error"), problem.Wrap(err)).JSONString(), http.StatusInternalServerError)
		return
	}

	http.Error(w, problem.New(problem.Title("internal server error")).JSONString(), http.StatusInternalServerError)
}

package handlers

import (
	"bacon-api/utils"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// db path
const (
	db = "db.json"
)

// TODO: Abstract sync into solo database handler
var guard sync.Mutex

func outputHeader(w http.ResponseWriter, s int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(s)
}

// Fails for checking errors in requests
func fails(e error, s string, w http.ResponseWriter) bool {
	if utils.Check(e, s) {
		outputHeader(w, http.StatusBadRequest)

		err := json.NewEncoder(w).Encode(s)

		if err != nil {
			log.Println("Handling failed!")
		}

		return true
	}

	return false
}

package handlers

import (
	"bacon-api/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
)

// GetMeeting handler
func GetMeeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	f, err := ioutil.ReadFile(db)

	if fails(err, "Failed to open data files", w) {
		return
	}

	m := gjson.GetBytes(f, id)

	if !m.Exists() {
		outputHeader(w, http.StatusBadRequest)

		err := json.NewEncoder(w).Encode("Failed to find meet")

		if err != nil {
			log.Println("Handling failed!")
		}

		return
	}

	meet := new(models.Meet)
	err = json.Unmarshal([]byte(m.String()), meet)

	if fails(err, "Failed to load data files", w) {
		return
	}

	outputHeader(w, http.StatusFound)

	err = json.NewEncoder(w).Encode(meet)

	if fails(err, "Failed to get new meeting", w) {
		return
	}
}

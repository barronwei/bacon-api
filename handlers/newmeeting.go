package handlers

import (
	"bacon-api/models"
	"bacon-api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/sjson"
)

type addMeetRequest struct {
	Name string
	When []int
}

// NewMeeting handler
func NewMeeting(w http.ResponseWriter, r *http.Request) {
	req := new(addMeetRequest)
	err := json.NewDecoder(r.Body).Decode(req)

	if fails(err, "Failed request", w) {
		return
	}

	name, when := req.Name, req.When
	team := models.People{}

	guard.Lock()
	defer guard.Unlock()

	f, err := ioutil.ReadFile(db)

	if fails(err, "Failed to open data files", w) {
		return
	}

	meetID := utils.NewID(f)
	m := models.NewMeet(meetID, name, team, when)

	s, err := sjson.SetBytes(f, meetID, m)

	if fails(err, "Failed to set new meeting", w) {
		return
	}

	if fails(err, "Failed to build a meeting", w) {
		return
	}

	err = ioutil.WriteFile(db, s, 0644)

	if fails(err, "Failed to add new meeting", w) {
		return
	}

	outputHeader(w, http.StatusOK)

	err = json.NewEncoder(w).Encode(meetID)

	if fails(err, "Failed to add new meeting", w) {
		return
	}
}

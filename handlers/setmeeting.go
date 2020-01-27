package handlers

import (
	"bacon-api/models"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type addUserRequest struct {
	ID   string
	PW   string
	Name string
	When []int
}

// SetMeeting handler
func SetMeeting(w http.ResponseWriter, r *http.Request) {
	req := new(addUserRequest)
	err := json.NewDecoder(r.Body).Decode(req)

	if fails(err, "Failed request", w) {
		return
	}

	id, name := req.ID, req.Name
	pw, when := req.PW, req.When

	u := uuid.New()

	h := sha256.New()
	pwHash := h.Sum([]byte(pw))

	user := models.NewUser(u.String(), name, pwHash, when)

	guard.Lock()
	defer guard.Unlock()

	f, err := ioutil.ReadFile(db)

	if fails(err, "Failed to open data files", w) {
		return
	}

	m := gjson.GetBytes(f, id)

	if !m.Exists() {
		outputHeader(w, http.StatusBadRequest)

		err = json.NewEncoder(w).Encode("Meet is unavailable")

		if err != nil {
			log.Println("Handling failed!")
		}

		return
	}

	c := gjson.GetBytes(f, id+".Team."+name)

	if c.Exists() {
		u := new(models.User)
		err = json.Unmarshal([]byte(c.String()), u)

		if fails(err, "Failed to handle user", w) {
			return
		}

		if !bytes.Equal(u.PW, user.PW) {
			outputHeader(w, http.StatusBadRequest)

			err = json.NewEncoder(w).Encode("Name already in use")

			if err != nil {
				log.Println("Handling failed!")
			}

			return
		}

		f, err = sjson.DeleteBytes(f, id+".Team."+name)

		if fails(err, "Failed to clear user", w) {
			return
		}
	}

	res, err := sjson.SetBytes(f, id+".Team."+name, user)

	if fails(err, "Failed to set new user", w) {
		return
	}

	err = ioutil.WriteFile(db, res, 0644)

	if fails(err, "Failed to add new user", w) {
		return
	}

	outputHeader(w, http.StatusCreated)

	err = json.NewEncoder(w).Encode(user)

	if fails(err, "Failed to set new user", w) {
		return
	}
}

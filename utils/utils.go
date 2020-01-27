package utils

import (
	"log"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"
)

// Check for checking errors
func Check(e error, s string) bool {
	if e != nil {
		log.Println(s)
		return true
	}

	return false
}

// NewID for creating a meet
func NewID(f []byte) string {
	m := uuid.New()
	r := gjson.GetBytes(f, m.String())

	if !r.Exists() {
		return m.String()
	}

	return NewID(f)
}

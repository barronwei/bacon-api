package models

import (
	"time"
)

// Span struct
type Span struct {
	Fst time.Time
	Snd time.Time
}

// Spans struct
type Spans []Span

// User struct
type User struct {
	ID   string
	Name string
	PW   []byte
	When []int
}

// People struct
type People map[string]User

// Meet struct
type Meet struct {
	ID   string
	Name string
	Team People
	When []int
}

// NewUser constructor
func NewUser(a string, b string, c []byte, d []int) *User {
	u := User{a, b, c, d}
	return &u
}

// NewMeet constructor
func NewMeet(a string, b string, c People, d []int) *Meet {
	m := Meet{a, b, c, d}
	return &m
}

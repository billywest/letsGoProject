package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord ...
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials ...
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail ...
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet ...
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User ...
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}

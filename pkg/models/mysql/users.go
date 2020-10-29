package mysql

import (
	"database/sql"

	"letsgo.net/snippetbox/pkg/models"
)

// UserModel ...
type UserModel struct {
	DB *sql.DB
}

// Insert ...
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate ...
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get ...
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}

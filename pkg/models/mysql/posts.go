package mysql

import (
	"database/sql"

	"github.com/turamant/Go-Blog/pkg/models"
)

type PostModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *PostModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// This will return a specific snippet based on its id.
func (m *PostModel) Get(id int) (*models.Post, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func (m *PostModel) Latest() ([]*models.Post, error) {
	return nil, nil
}
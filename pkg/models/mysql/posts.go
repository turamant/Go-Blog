package mysql

import (
	"database/sql"
	"errors"

	"github.com/turamant/Go-Blog/pkg/models"
)

type PostModel struct {
	DB *sql.DB
}

// This will insert a new post into the database.
func (m *PostModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO posts (title, content, created, expires) 
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// This will return a specific post based on its id.
func (m *PostModel) Get(id int) (*models.Post, error) {
	stmt := `SELECT id, title, content, created, expires FROM posts
	WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.Post{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// This will return the 10 most recently created posts.
func (m *PostModel) Latest() ([]*models.Post, error) {
	stmt := `SELECT id, title, content, created, expires FROM posts 
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []*models.Post{}
	for rows.Next() {
		s := &models.Post{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		posts = append(posts, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m *PostModel) All() ([]*models.Post, error) {
	stmt := `SELECT id, title, content, created, expires FROM posts`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []*models.Post{}
	for rows.Next() {
		s := &models.Post{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		posts = append(posts, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}


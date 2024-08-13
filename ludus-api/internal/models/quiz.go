package models

import (
	"database/sql"
	"time"
)

type Quiz struct {
	ID		int
	Title	string
	Content string
	Created time.Time
}

type QuizModel struct {
	DB *sql.DB
}

func (m *QuizModel) Insert(title string, content string) (int, error) {
	return 0, nil
}

func (m *QuizModel) Get(id int) (Quiz, error) {
	return Quiz{}, nil
}

func (m *QuizModel) Latest() ([]Quiz, error) {
	return nil, nil
}
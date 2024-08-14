package models

import (
	"database/sql"
	"errors"
	"time"
)

type Quiz struct {
	ID	int
	Title	string
	Content string
	Created time.Time
}

type QuizModel struct {
	DB *sql.DB
}

func (m *QuizModel) Insert(title string, content string) (int, error) {
	stmt := `INSERT INTO quiz (title, content, created)
			 VALUES(?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *QuizModel) Get(id int) (Quiz, error) {
	var q Quiz
	stmt := `SELECT id, title, content, created
			 FROM quiz WHERE id = ?`
	
	err := m.DB.QueryRow(stmt, id).Scan(&q.ID, &q.Title, &q.Content, &q.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Quiz{}, ErrNoRecord
		}
		return Quiz{}, err
	}

	return q, nil
}

func (m *QuizModel) Latest() ([]Quiz, error) {
	stmt := `SELECT id, title, content, created
			 FROM quiz ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var quizzes []Quiz

	for rows.Next() {
		var q Quiz
		err := rows.Scan(&q.ID, &q.Title, &q.Content, &q.Created)
		if err != nil {
			return nil, err
		}
		quizzes = append(quizzes, q)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return quizzes, nil
}
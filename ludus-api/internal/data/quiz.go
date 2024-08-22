package data

import (
	"time"

	"github.com/KaliszS/Ludus/internal/validator"
)

type Quiz struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

func ValidateQuiz(v *validator.Validator, quiz *Quiz) {
	v.Check(quiz.Title != "", "title", "title must be provided")
	v.Check(len(quiz.Title) <= 100, "title", "title must not be more than 100 bytes long")
}

package main

import (
	"errors"
    "fmt"
    "net/http"
    "strconv"

	"github.com/KaliszS/Ludus/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {	
	quizzes, err := app.quiz.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, quiz := range quizzes {
		fmt.Fprintf(w, "%+v\n", quiz)
	}
}

func (app *application) quizView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	quiz, err := app.quiz.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", quiz)
}

func (app *application) quizCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Quiz created"))
}

func (app *application) quizCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "Example Quiz"
	content := "This is an example quiz"

	id, err := app.quiz.Insert(title, content)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/quiz/%d", id), http.StatusSeeOther)
}
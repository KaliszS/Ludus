package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KaliszS/Ludus/internal/data"
	"github.com/KaliszS/Ludus/internal/validator"
)

func (app *application) quizViewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	quiz := data.Quiz{
		ID:      id,
		Title:   "Example Quiz",
		Content: "This is an example quiz",
		Created: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"quiz": quiz}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	// quiz, err := app.quiz.Get(int(id))
	// if err != nil {
	// 	if errors.Is(err, models.ErrNoRecord) {
	// 		http.NotFound(w, r)
	// 	} else {
	// 		app.serverError(w, r, err)
	// 	}
	// 	return
	// }

	// fmt.Fprintf(w, "%+v", quiz)
}

func (app *application) quizCreateHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	quiz := &data.Quiz{
		Title:   input.Title,
		Content: input.Content,
		Created: time.Now(),
	}

	v := validator.New()

	if data.ValidateQuiz(v, quiz); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)

	// title := "Example Quiz"
	// content := "This is an example quiz"

	// id, err := app.quiz.Insert(title, content)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }

	// http.Redirect(w, r, fmt.Sprintf("/quiz/%d", id), http.StatusSeeOther)
}

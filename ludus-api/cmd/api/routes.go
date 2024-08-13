package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /quiz/view/{id}", app.quizView)
	mux.HandleFunc("GET /quiz/create", app.quizCreate)
	mux.HandleFunc("POST /quiz/create", app.quizCreatePost)

	return mux
}
package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /quiz/{id}", app.quizView)
	mux.HandleFunc("GET /quiz/create", app.quizCreate)
	mux.HandleFunc("POST /quiz", app.quizCreatePost)

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, commonHeaders, app.sessionManager.LoadAndSave)

	return standardMiddleware.Then(mux)
}
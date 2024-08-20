package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1", app.home)
	router.HandlerFunc(http.MethodGet, "/v1/quiz/:id", app.quizView)
	router.HandlerFunc(http.MethodPost, "/v1/quiz", app.quizCreate)

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, commonHeaders, app.sessionManager.LoadAndSave)

	return standardMiddleware.Then(router)
}

package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello, World!"))
}

func (app *application) quizView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display quiz with ID %d", id)
}

func (app *application) quizCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Quiz created"))
}

func (app *application) quizCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Quiz created using post"))
}
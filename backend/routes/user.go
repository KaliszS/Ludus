package routes

import (
	"net/http"

	"github.com/google/uuid"

	"ludus/services"
)

type UserHandler struct {
	userService services.UserCrud
}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idInput := r.URL.Query().Get("id")
	id, err := uuid.Parse(idInput)
	if err != nil {
		w.Write([]byte("Invalid ID"))
		return
	}
	name := r.URL.Query().Get("name")

	user := handler.userService.GetUser(id, name)
	w.Write([]byte("User: " + user.Name))
}
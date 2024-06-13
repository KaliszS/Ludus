package services

import (
	"log"

	"github.com/google/uuid"

	"ludus/models"
)


type UserCrud struct {}

func (handler *UserCrud) GetUser(id uuid.UUID, name string) models.User {
	log.Println("Getting user with id", id, "and name", name)
	return models.User{
		ID:   id,
		Name: name,
	}
}
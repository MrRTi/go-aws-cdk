package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("Bad request: username and passord are required")
	}

	userExists, err := api.dbStore.UserExists(event.Username)

	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	if userExists {
		return fmt.Errorf("User with that username already exists")
	}

	err = api.dbStore.InsertUser(event)

	if err != nil {
		return fmt.Errorf("Error when creating user: %w", err)
	}

	return nil
}

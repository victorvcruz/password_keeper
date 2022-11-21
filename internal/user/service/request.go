package service

type UserRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	MasterPassword string `json:"masterPassword"`
}

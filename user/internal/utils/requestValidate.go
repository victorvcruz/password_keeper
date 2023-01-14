package utils

import "github.com/go-playground/validator/v10"

func RequestUserValidate(err error) string {
	for _, err := range err.(validator.ValidationErrors) {
		if err.Namespace() == "UserRequest.Name" && err.Tag() == "gte" {
			return "Short name"
		}
		if err.Namespace() == "UserRequest.Name" && err.Tag() == "lte" {
			return "Long name"
		}
		if err.Namespace() == "UserRequest.Email" {
			return "Invalid email"
		}
		if err.Namespace() == "UserRequest.MasterPassword" && err.Tag() == "gte" {
			return "Short password"
		}
		if err.Namespace() == "UserRequest.MasterPassword" && err.Tag() == "lte" {
			return "Long password"
		}
	}
	return ""
}

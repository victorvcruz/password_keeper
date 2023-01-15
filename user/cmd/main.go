package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
	"user.com/cmd/api"
	"user.com/cmd/api/handlers"
	auth "user.com/internal/auth"
	"user.com/internal/crypto"
	dbmanager "user.com/internal/platform/database"
	"user.com/internal/user"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	database, err := dbmanager.NewDatabase().Connect()
	if err != nil {
		log.Fatalf("[CONNECT DATABASE FAIL]: %s", err.Error())
	}

	authService := auth.NewAuthService()

	crypto := crypto.NewCrypto(os.Getenv("CRYPTO_SECRET"))
	validate := validator.New()
	userRepository := user.NewUserRepository(database)
	userService := user.NewUserService(userRepository, crypto)

	userHandler := handlers.NewUserHandler(userService, authService, validate)

	err = api.New(userHandler)
	if err != nil {
		log.Fatalf("[START SERVER  FAIL]: %s", err.Error())
	}
}

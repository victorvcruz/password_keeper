package main

import (
	"auth.com/cmd/api"
	"auth.com/cmd/api/handlers"
	"auth.com/internal/auth"
	"auth.com/internal/crypto"
	dbmanager "auth.com/internal/platform/database"
	"auth.com/internal/token"
	"auth.com/internal/user"
	"auth.com/internal/utils/authorization"
	"github.com/joho/godotenv"
	"log"
)

const Service = "AUTH"

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

	authorization := authorization.NewAuthorization(Service)

	userService := user.NewUserService(authorization)
	crypto := crypto.NewCrypto()
	token := token.NewTokenService()

	authRepository := auth.NewAuthRepository(database)

	authService := auth.NewAuthService(database, authRepository, userService, crypto, token)

	_, err = authService.RegisterService(&auth.Register{Service: Service})
	if err != nil {
		log.Fatalf("Failed to register service %s", err.Error())
	}

	authHandler := handlers.NewAuthHandler(authService)

	err = api.New(authHandler)
	if err != nil {
		log.Fatalf("[START SERVER FAIL]: %s", err.Error())
	}
}

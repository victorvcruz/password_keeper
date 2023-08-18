package main

import (
	"github.com/joho/godotenv"
	"log"
	"vault.com/cmd/api"
	"vault.com/cmd/api/handlers"
	dbmanager "vault.com/internal/platform/database"
	"vault.com/internal/vault"
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

	vaultRepository := vault.NewVaultRepository(database)

	vaultService := vault.NewVaultService(database, vaultRepository)

	vaultHandler := handlers.NewVaultHandler(vaultService)

	err = api.New(vaultHandler)
	if err != nil {
		log.Fatalf("[START SERVER FAIL]: %s", err.Error())
	}
}

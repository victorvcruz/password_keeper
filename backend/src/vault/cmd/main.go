package main

import (
	"github.com/joho/godotenv"
	"log"
	"vault.com/cmd/api"
	"vault.com/cmd/api/handlers"
	"vault.com/internal/auth"
	"vault.com/internal/folder"
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

	folderRepository := folder.NewFolderRepository(database)

	folderService := folder.NewFolderService(database, folderRepository)

	authService := auth.NewAuthService()

	vaultHandler := handlers.NewVaultHandler(vaultService, folderService, authService)

	folderHandler := handlers.NewFolderHandler(folderService, authService)

	err = api.New(vaultHandler, folderHandler)
	if err != nil {
		log.Fatalf("[START SERVER FAIL]: %s", err.Error())
	}
}

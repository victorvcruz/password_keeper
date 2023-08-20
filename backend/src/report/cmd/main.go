package main

import (
	"github.com/joho/godotenv"
	"log"
	"report.com/cmd/api"
	"report.com/cmd/api/handlers"
	dbmanager "report.com/internal/platform/database"
	"report.com/internal/report"
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

	reportRepository := report.NewReportRepository(database)

	reportService := report.NewReportService(reportRepository)

	reportHandler := handlers.NewReportHandler(reportService)

	err = api.New(reportHandler)
	if err != nil {
		log.Fatalf("[START SERVER FAIL]: %s", err.Error())
	}
}

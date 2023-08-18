package main

import (
	"github.com/joho/godotenv"
	"log"
	"report.com/cmd/api"
	"report.com/cmd/api/handlers"
	"report.com/internal/auth"
	dbmanager "report.com/internal/platform/database"
	"report.com/internal/report"
	"report.com/pkg/authorization"
)

const Service = "REPORT"

func init() {
	authorization.Setup(Service)
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

	err = api.New(reportHandler, auth.NewAuthService(Service))
	if err != nil {
		log.Fatalf("[START SERVER FAIL]: %s", err.Error())
	}
}

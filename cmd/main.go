package main

import (
	notestracker "NotesTracker"
	"NotesTracker/database"
	"NotesTracker/pkg/handler"
	"NotesTracker/pkg/repository"
	"NotesTracker/pkg/service"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := database.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}

	db, err := database.NewDBConnection(dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	h := handler.NewHandler(service)

	router := h.InitRoutes()

	server := notestracker.Server{}
	if err := server.Run("8080", router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

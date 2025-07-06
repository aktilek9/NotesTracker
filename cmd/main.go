package main

import (
	"NotesTracker/pkg/handler"
	"NotesTracker/pkg/repository"
	"NotesTracker/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	srv := service.NewService(repo)
	h := handler.NewHandler(srv)

	router := h.InitRoutes()
	router.Run(":8080")
}

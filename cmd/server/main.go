package main

import (
	"log"
	"net/http"

	"github.com/tflorimo/habit-tracker/internal/handler"

	"github.com/tflorimo/habit-tracker/internal/config"
)

func main() {
	cfg := config.NewConfig()
	http.HandleFunc("/", handler.HandleRoot)
	http.HandleFunc("/habits", handler.NewHabitHandler().HandleHabitRoutes)
	http.HandleFunc("/habits/tracking", handler.HandleTrackingRoutes)
	http.HandleFunc("/habits/category", handler.HandleCategoryRoutes)

	log.Printf("Server started on port %s", cfg.HttpPort)
	err := http.ListenAndServe(":"+cfg.HttpPort, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

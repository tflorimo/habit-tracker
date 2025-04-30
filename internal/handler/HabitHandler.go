package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tflorimo/habit-tracker/internal/model"
	"github.com/tflorimo/habit-tracker/internal/service"
)

type HabitHandler struct {
	service *service.HabitService
}

func NewHabitHandler() *HabitHandler {
	return &HabitHandler{}
}

func (hh *HabitHandler) HandleHabitRoutes(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/habits/")
	if id == "" {
		// no id provided, return 400
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No habit ID provided"))
		return
	}
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
	case "POST":
		hh.PostHabit(w, r)
	case "PUT":

		w.WriteHeader(http.StatusOK)
	case "DELETE":

		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func (hh *HabitHandler) PostHabit(w http.ResponseWriter, r *http.Request) {
	var habit model.Habit
	if err := json.NewDecoder(r.Body).Decode(&habit); err != nil {
		http.Error(w, "The payload is invalid", http.StatusBadRequest)
		return
	}

	err := hh.service.CreateHabit(&habit)
	if err != nil {
		http.Error(w, "There was an error creating the habit", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("The habit was created successfully"))

}

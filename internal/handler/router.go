package handler

import (
	"net/http"
	"strings"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Habit Tracker API!"))
}

func HandleCategoryRoutes(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/habits/category/")
	if id == "" {
		// no id provided, return 400
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No category ID provided"))
		return
	}
	switch r.Method {
	case "GET":

		w.WriteHeader(http.StatusOK)
	case "POST":

		w.WriteHeader(http.StatusCreated)
	case "PUT":

		w.WriteHeader(http.StatusOK)
	case "DELETE":

		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func HandleTrackingRoutes(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/habits/tracking/")
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

		w.WriteHeader(http.StatusCreated)
	case "PUT":

		w.WriteHeader(http.StatusOK)
	case "DELETE":

		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

package repository

import (
	"database/sql"

	"github.com/tflorimo/habit-tracker/internal/model"
)

type HabitRepository struct {
	db *sql.DB
}

func NewHabitRepository(db *sql.DB) *HabitRepository {
	return &HabitRepository{db: db}
}

func (r *HabitRepository) Insert(h *model.Habit) error {
	query := "INSERT INTO habits (name, categoryId) VALUES (?, ?)"
	_, err := r.db.Exec(query, h.Name, h.CategoryID)
	return err
}

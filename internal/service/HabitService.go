package service

import (
	"github.com/tflorimo/habit-tracker/internal/model"
	"github.com/tflorimo/habit-tracker/internal/repository"
)

type HabitService struct {
	repo *repository.HabitRepository
}

func NewHabitService() *HabitService {
	return &HabitService{}
}

func (s *HabitService) CreateHabit(h *model.Habit) error {
	return s.repo.Insert(h)
}

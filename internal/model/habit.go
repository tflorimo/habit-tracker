package model

import "time"

type Habit struct {
	ID          int
	Name        string
	Description string
	CategoryID  int // FK for HabitCategory
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewHabit(ID int, Name string, Description string, CategoryID int) *Habit {
	h := Habit{ID: ID, Name: Name, Description: Description, CategoryID: CategoryID}

	return &h
}

package model

import "time"

type Habit struct {
	ID         int
	Name       string
	CategoryID int // FK for HabitCategory
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func NewHabit(ID int, Name string, CategoryID int) *Habit {
	h := Habit{ID: ID, Name: Name, CategoryID: CategoryID}
	return &h
}

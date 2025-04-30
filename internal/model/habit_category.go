package model

import "time"

type HabitCategory struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewHabitCategory(ID int, Name string) *HabitCategory {
	hc := HabitCategory{ID: ID, Name: Name}
	hc.CreatedAt = time.Now()

	return &hc
}

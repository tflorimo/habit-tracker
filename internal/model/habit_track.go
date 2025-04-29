package model

import "time"

type HabitTrack struct {
	ID        int
	HabitID   int       // FK for Habit
	CreatedAt time.Time // if it was inserted we assume it was "checked" as "completed"
	UpdatedAt time.Time
}

func newHabitTrack(ID int, HabitID int) *HabitTrack {
	hTrack := HabitTrack{ID: ID, HabitID: HabitID}
	hTrack.CreatedAt = time.Now()

	return &hTrack
}

func setUpdatedNow(hTrack *HabitTrack) {
	hTrack.UpdatedAt = time.Now()
}

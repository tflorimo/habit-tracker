package model

type Habit struct {
	ID          int
	Name        string
	Description string
	CategoryID  int
}

func newHabit(ID int, Name string, Description string, CategoryID int) *Habit {
	h := Habit{ID: ID, Name: Name, Description: Description, CategoryID: CategoryID}

	return &h
}

package domain

import "github.com/google/uuid"

type MuscleGroup struct {
	ID   uuid.UUID
	Name string
}

func NewMuscleGroup(ID uuid.UUID, name string) *MuscleGroup {
	return &MuscleGroup{ID: ID, Name: name}
}

package domain

import "github.com/google/uuid"

type Exercise struct {
	ID          uuid.UUID
	Name        string
	MuscleGroup string
}

func NewExercise(name string, mg string) (Exercise, error) {
	return Exercise{
		ID:          uuid.New(),
		Name:        name,
		MuscleGroup: mg,
	}, nil
}

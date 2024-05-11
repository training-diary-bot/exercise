package dbview

import (
	"exercise/internal/domain"
	"github.com/google/uuid"
)

type Exercise struct {
	ID          uuid.UUID
	Name        string
	MuscleGroup string
}

type MuscleGroups struct {
	ID   uuid.UUID
	Name string
}

func ExerciseToDomain(ex Exercise) (domain.Exercise, error) {
	return domain.Exercise{
		ID:          ex.ID,
		Name:        ex.Name,
		MuscleGroup: ex.MuscleGroup,
	}, nil
}

//func MuscleGroupsToDomain(mg MuscleGroups) (domain.MuscleGroups, error) {
//	return domain.MuscleGroups{
//		ID:   mg.ID,
//		Name: mg.Name,
//	}, nil
//}

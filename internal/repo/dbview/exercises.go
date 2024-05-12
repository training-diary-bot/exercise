package dbview

import (
	"exercise/internal/domain"
	"github.com/google/uuid"
)

type Exercise struct {
	ID            uuid.UUID
	Name          string
	MuscleGroupID string
}

func (ex Exercise) ToDomain() domain.Exercise {
	return domain.Exercise{
		ID:            ex.ID,
		Name:          ex.Name,
		MuscleGroupID: ex.MuscleGroupID,
	}
}

type MuscleGroup struct {
	ID   uuid.UUID
	Name string
}

func (mg MuscleGroup) ToDomain() domain.MuscleGroup {
	return domain.MuscleGroup{
		ID:   mg.ID,
		Name: mg.Name,
	}
}

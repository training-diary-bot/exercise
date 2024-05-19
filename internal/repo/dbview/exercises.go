package dbview

import (
	"exercise/internal/domain"
	"github.com/google/uuid"
)

type Exercise struct {
	ID            uuid.UUID `db:"id"`
	Name          string    `db:"name"`
	MuscleGroupID string    `db:"muscle_group_id"`
}

func (ex Exercise) ToDomain() domain.Exercise {
	return domain.Exercise{
		ID:            ex.ID,
		Name:          ex.Name,
		MuscleGroupID: ex.MuscleGroupID,
	}
}

type MuscleGroup struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func (mg MuscleGroup) ToDomain() domain.MuscleGroup {
	return domain.MuscleGroup{
		ID:   mg.ID,
		Name: mg.Name,
	}
}

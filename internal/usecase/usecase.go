package usecase

import (
	"context"
	"exercise/internal/domain"
	"github.com/google/uuid"
)

type ISQLRepo interface {
	CreateCustomExercise(ctx context.Context, userID uuid.UUID, name string) (uuid.UUID, error)
	GetExercisesByMuscleGroup(ctx context.Context, userID uuid.UUID, muscleGroupID uint64) ([]domain.Exercise, error)
	GetMuscleGroups(ctx context.Context) ([]string, error)
	GetCustomExercises(ctx context.Context, userID uuid.UUID) (domain.Exercise, error)
}

type UC struct {
	sqlRepo ISQLRepo
}

func NewUC(repo ISQLRepo) UC {
	return UC{
		sqlRepo: repo,
	}
}

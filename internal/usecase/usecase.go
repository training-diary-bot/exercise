package usecase

import (
	"context"
	"exercise/internal/domain"
	"github.com/google/uuid"
)

type iSQLRepo interface {
	GetMuscleGroups(ctx context.Context) ([]domain.MuscleGroup, error)
	GetMuscleGroupsByIDs(ctx context.Context, ids []uuid.UUID) ([]domain.MuscleGroup, error)
	GetExercisesByMuscleGroupID(ctx context.Context, id uuid.UUID) ([]domain.Exercise, error)
	GetExercisesByIDs(ctx context.Context, ids []uuid.UUID) ([]domain.Exercise, error)
}

type UC struct {
	sqlRepo iSQLRepo
}

func NewUC(repo iSQLRepo) UC {
	return UC{
		sqlRepo: repo,
	}
}

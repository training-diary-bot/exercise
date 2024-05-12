package usecase

import (
	"context"
	"exercise/internal/domain"
	"github.com/google/uuid"
)

func (uc UC) GetMuscleGroups(ctx context.Context) ([]domain.MuscleGroup, error) {
	return uc.sqlRepo.GetMuscleGroups(ctx)
}

func (uc UC) GetMuscleGroupsByIDs(ctx context.Context, ids []uuid.UUID) ([]domain.MuscleGroup, error) {
	return uc.sqlRepo.GetMuscleGroupsByIDs(ctx, ids)
}

func (uc UC) GetExercisesByMuscleGroupID(ctx context.Context, id uuid.UUID) ([]domain.Exercise, error) {
	return uc.sqlRepo.GetExercisesByMuscleGroupID(ctx, id)
}

func (uc UC) GetExercisesByIDs(ctx context.Context, ids []uuid.UUID) ([]domain.Exercise, error) {
	return uc.sqlRepo.GetExercisesByIDs(ctx, ids)
}

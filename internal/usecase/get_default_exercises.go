package usecase

import (
	"context"
	"exercise/internal/domain"
	"github.com/google/uuid"
)

func (uc UC) GetExercisesByMuscleGroup(ctx context.Context, userID uuid.UUID, muscleGroupID uint64) (domain.Exercise, error) {
	return uc.sqlRepo.GetExercisesByMuscleGroup(ctx, userID, muscleGroupID)
}

package usecase

import (
	"context"
	"exercise/internal/domain"
	"github.com/google/uuid"
)

func (uc UC) GetCustomExercises(ctx context.Context, userID uuid.UUID) (domain.Exercise, error) {
	return uc.sqlRepo.GetCustomExercises(ctx, userID)
}

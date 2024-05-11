package usecase

import (
	"context"
	"github.com/google/uuid"
)

func (uc UC) CreateCustomExercise(ctx context.Context, userID uuid.UUID, name string) (uuid.UUID, error) {
	ex, err := uc.sqlRepo.CreateCustomExercise(ctx, userID, name)
	if err != nil {
		return uuid.UUID{}, nil
	}

	return ex, nil
}

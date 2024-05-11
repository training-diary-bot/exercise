package usecase

import (
	"context"
)

func (uc UC) GetMuscleGroups(ctx context.Context) ([]string, error) {
	return uc.sqlRepo.GetMuscleGroups(ctx)
}

package grpc

import (
	"exercise/internal/domain"
	commonv1 "github.com/training-diary-bot/apis/genproto/go/common/v1"
	pb "github.com/training-diary-bot/apis/genproto/go/exercise/v1"
)

func ExerciseToProto(ex []domain.Exercise) *pb.GetCustomExercisesResponse {
	res := make([]*commonv1.Exercise, 0, len(ex))

	for _, val := range ex {
		res = append(res, &commonv1.Exercise{
			Id:   val.ID.String(),
			Name: val.Name,
		})
	}

	return &pb.GetCustomExercisesResponse{
		Exercises: res,
	}
}

func MuscleGroupToProto(mg []domain.MuscleGroup) *pb.GetMuscleGroupsResponse {
	res := make([]*commonv1.MuscleGroup, 0, len(mg))

	for _, val := range mg {
		res = append(res, &commonv1.MuscleGroup{
			Id:   val.ID.String(),
			Name: val.Name,
		})
	}

	return &pb.GetMuscleGroupsResponse{
		MuscleGroups: res,
	}
}

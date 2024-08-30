package grpc

import (
	"exercise/internal/domain"
	pb "github.com/training-diary-bot/apis/genproto/go/exercise/v1"
)

func MuscleGroupsToProto(mg []domain.MuscleGroup) []*pb.MuscleGroupEntity {
	res := make([]*pb.MuscleGroupEntity, 0, len(mg))

	for _, val := range mg {
		res = append(res, &pb.MuscleGroupEntity{
			Id:   val.ID.String(),
			Name: val.Name,
		})
	}

	return res
}

func ExercisesToProto(exs []domain.Exercise) []*pb.ExerciseEntity {
	res := make([]*pb.ExerciseEntity, 0, len(exs))

	for _, val := range exs {
		res = append(res, &pb.ExerciseEntity{
			Id:            val.ID.String(),
			Name:          val.Name,
			MuscleGroupId: val.MuscleGroupID,
		})
	}

	return res
}

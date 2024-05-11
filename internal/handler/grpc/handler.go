package grpc

import (
	"context"
	"exercise/internal/domain"
	"exercise/pkg/grpc"
	"exercise/pkg/logger"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	pb "github.com/training-diary-bot/apis/genproto/go/exercise/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IUseCase interface {
	CreateCustomExercise(ctx context.Context, userID uuid.UUID, name string) (uuid.UUID, error)
	GetExercisesByMuscleGroup(ctx context.Context, userID uuid.UUID, muscleGroupID uint64) ([]domain.Exercise, error)
	GetMuscleGroups(ctx context.Context) ([]domain.MuscleGroup, error)
	GetCustomExercises(ctx context.Context, userID uuid.UUID) ([]domain.Exercise, error)
}

type Handler struct {
	uc IUseCase
}

func NewHandler(uc IUseCase) Handler {
	return Handler{
		uc: uc,
	}
}

func (h Handler) Register(s *grpc.Server) {
	pb.RegisterExerciseServer(s.GRPCServer(), h)
}

func (h Handler) CreateCustomExercise(ctx context.Context, req *pb.CreateCustomExerciseRequest) (*pb.CreateCustomExerciseResponse, error) {
	log := logger.New()
	log.Info().Msgf("new grpc request: create custom exercise: exercise name: %s", req.Name)

	userID, err := uuid.Parse(req.UserId)
	id, err := h.uc.CreateCustomExercise(ctx, userID, req.Name)
	if err != nil {
		log.Error().Err(err).Msg("failed to create exercise")
	}

	return &pb.CreateCustomExerciseResponse{
		ExerciseId: id.String(),
	}, nil
}

func (h Handler) GetExercisesByMuscleGroup(ctx context.Context, req *pb.GetExercisesByMuscleGroupRequest) (*pb.GetExercisesByMuscleGroupResponse, error) {
	log := logger.New()
	log.Info().Msgf("new grpc request: get exercises by muscle group: userID: %d", req.UserId)

	return nil, nil
}

func (h Handler) GetMuscleGroups(ctx context.Context, empty *emptypb.Empty) (*pb.GetMuscleGroupsResponse, error) {
	log := logger.New()
	log.Info().Msgf("new grpc request: get exercises by muscle groups")

	mg, err := h.uc.GetMuscleGroups(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to create exercise")
		return nil, status.Error(codes.Internal, "cant get muscle group")
	}

	return MuscleGroupToProto(mg), nil

}
func (h Handler) GetCustomExercises(ctx context.Context, req *pb.GetCustomExercisesRequest) (*pb.GetCustomExercisesResponse, error) {
	log := logger.New()
	log.Info().Msgf("new grpc request: get custom exercise: userID: %d", req.UserId)

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "not valid user ID")
	}

	ce, err := h.uc.GetCustomExercises(ctx, userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to create exercise")
		return nil, status.Error(codes.Internal, "cant get muscle group")
	}

	return ExerciseToProto(ce), nil
}

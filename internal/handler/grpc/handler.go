package grpc

import (
	"context"
	"exercise/internal/domain"
	"exercise/pkg/grpc"
	"exercise/pkg/logger"
	"github.com/google/uuid"
	pb "github.com/training-diary-bot/apis/genproto/go/exercise/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type iUseCase interface {
	GetMuscleGroups(ctx context.Context) ([]domain.MuscleGroup, error)
	GetMuscleGroupsByIDs(ctx context.Context, ids []uuid.UUID) ([]domain.MuscleGroup, error)
	GetExercisesByMuscleGroupID(ctx context.Context, id uuid.UUID) ([]domain.Exercise, error)
	GetExercisesByIDs(ctx context.Context, ids []uuid.UUID) ([]domain.Exercise, error)
}

type Handler struct {
	uc iUseCase
}

func NewHandler(uc iUseCase) Handler {
	return Handler{
		uc: uc,
	}
}

func (h Handler) Register(s *grpc.Server) {
	pb.RegisterExerciseServer(s.GRPCServer(), h)
}

func (h Handler) GetExercisesByIDs(ctx context.Context, req *pb.GetExercisesByIDsRequest) (*pb.GetExercisesByIDsResponse, error) {
	log := logger.New()
	log.Info().Msgf("new grpc request: get exercises by ids")

	ids := make([]uuid.UUID, 0, len(req.Ids))
	for _, id := range req.Ids {
		uuidID, err := uuid.Parse(id)
		if err != nil {
			return nil, status.Error(codes.Internal, "cant parse uuid")
		}

		ids = append(ids, uuidID)
	}

	exs, err := h.uc.GetExercisesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("failed to get exercises by ids")
		return nil, status.Error(codes.Internal, "cant get exercises by ids")
	}

	return &pb.GetExercisesByIDsResponse{
		Exercises: ExercisesToProto(exs),
	}, nil
}

func (h Handler) GetMuscleGroupsByIDs(ctx context.Context, req *pb.GetMuscleGroupsByIDsRequest) (*pb.GetMuscleGroupsByIDsResponse, error) {
	log := logger.New()
	log.Info().Msgf("new grpc request: get muscle groups by ids")

	ids := make([]uuid.UUID, 0, len(req.Ids))
	for _, id := range req.Ids {
		uuidID, err := uuid.Parse(id)
		if err != nil {
			return nil, status.Error(codes.Internal, "cant parse uuid")
		}

		ids = append(ids, uuidID)
	}

	mgs, err := h.uc.GetMuscleGroupsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("failed to get muscle group by ids")
		return nil, status.Error(codes.Internal, "cant get muscle group by ids")
	}

	return &pb.GetMuscleGroupsByIDsResponse{
		MuscleGroups: MuscleGroupsToProto(mgs),
	}, nil
}

func (h Handler) GetMuscleGroups(ctx context.Context, req *emptypb.Empty) (*pb.GetMuscleGroupsResponse, error) {
	log := logger.New()
	log.Info().Msgf("new grpc request: get muscle groups")

	mgs, err := h.uc.GetMuscleGroups(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get muscle group")
		return nil, status.Error(codes.Internal, "cant get muscle group")
	}

	return &pb.GetMuscleGroupsResponse{
		MuscleGroups: MuscleGroupsToProto(mgs),
	}, nil
}

func (h Handler) GetExercisesByMuscleGroupID(ctx context.Context, req *pb.GetExercisesByMuscleGroupIDRequest) (*pb.GetExercisesByMuscleGroupIDResponse, error) {
	log := logger.New()
	log.Info().Msgf("new grpc request: get exercises by muscle group id")

	id, err := uuid.Parse(req.MuscleGroupId)
	if err != nil {
		return nil, status.Error(codes.Internal, "cant parse muscle group id")
	}

	exs, err := h.uc.GetExercisesByMuscleGroupID(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("failed to get exercises by muscle group id")
		return nil, status.Error(codes.Internal, "cant get exercises by muscle group id")
	}

	return &pb.GetExercisesByMuscleGroupIDResponse{
		Exercises: ExercisesToProto(exs),
	}, nil
}

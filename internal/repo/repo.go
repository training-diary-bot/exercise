package repo

import (
	"context"
	"exercise/internal/domain"
	"exercise/internal/repo/dbview"
	"exercise/pkg/database"
	"exercise/pkg/logger"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Repo struct {
	db database.IDatabase
}

func NewRepo(db database.IDatabase) Repo {
	return Repo{
		db: db,
	}
}

func (r Repo) GetMuscleGroups(ctx context.Context) ([]domain.MuscleGroup, error) {
	log := logger.New()

	req, args, err := sq.
		Select(dbview.GetMuscleGroupsAllFields()...).
		From(dbview.MuscleGroupTable).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get muscle groups")
	}

	log.Debug().Msgf("query: %s %q", req, args)

	rows, err := r.db.Query(ctx, req, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query exec")
	}

	var res []domain.MuscleGroup

	for rows.Next() {
		var muscleGroupsView dbview.MuscleGroup

		if err = pgxscan.ScanRow(&muscleGroupsView, rows); err != nil {
			log.Error().Msg(errors.Wrap(err, "scan err").Error())
			return nil, errors.Wrap(err, "scan err")
		}

		res = append(res, muscleGroupsView.ToDomain())
	}

	return res, nil
}

func (r Repo) GetMuscleGroupsByIDs(ctx context.Context, ids []uuid.UUID) ([]domain.MuscleGroup, error) {
	log := logger.New()

	req, args, err := sq.
		Select(dbview.GetMuscleGroupsAllFields()...).
		From(dbview.MuscleGroupTable).
		Where(sq.Eq{
			dbview.MuscleGroupsIDField: ids,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get muscle groups by ids")
	}

	log.Debug().Msgf("query: %s %q", req, args)

	rows, err := r.db.Query(ctx, req, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query exec")
	}

	res := make([]domain.MuscleGroup, 0, len(ids))

	for rows.Next() {
		var muscleGroupsView dbview.MuscleGroup

		if err = pgxscan.ScanRow(&muscleGroupsView, rows); err != nil {
			log.Error().Msg(errors.Wrap(err, "scan err").Error())
			return nil, errors.Wrap(err, "scan err")
		}

		res = append(res, muscleGroupsView.ToDomain())
	}

	return res, nil
}

func (r Repo) GetExercisesByMuscleGroupID(ctx context.Context, id uuid.UUID) ([]domain.Exercise, error) {
	log := logger.New()

	req, args, err := sq.
		Select(dbview.GetExercisesAllFields()...).
		From(dbview.ExercisesTable).
		Where(sq.Eq{
			dbview.ExercisesMuscleGroupField: id,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get exercise by muscle group id")
	}

	log.Debug().Msgf("query: %s %q", req, args)

	rows, err := r.db.Query(ctx, req, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query exec")
	}

	var res []domain.Exercise

	for rows.Next() {
		var exView dbview.Exercise

		if err = pgxscan.ScanRow(&exView, rows); err != nil {
			log.Error().Msg(errors.Wrap(err, "scan err").Error())
			return nil, errors.Wrap(err, "scan err")
		}

		res = append(res, exView.ToDomain())
	}

	return res, nil
}

func (r Repo) GetExercisesByIDs(ctx context.Context, ids []uuid.UUID) ([]domain.Exercise, error) {
	log := logger.New()

	req, args, err := sq.
		Select(dbview.GetExercisesAllFields()...).
		From(dbview.ExercisesTable).
		Where(sq.Eq{
			dbview.ExercisesIDField: ids,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get exercise by ids")
	}

	log.Debug().Msgf("query: %s %q", req, args)

	rows, err := r.db.Query(ctx, req, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query exec")
	}

	var res []domain.Exercise

	for rows.Next() {
		var exView dbview.Exercise

		if err = pgxscan.ScanRow(&exView, rows); err != nil {
			log.Error().Msg(errors.Wrap(err, "scan err").Error())
			return nil, errors.Wrap(err, "scan err")
		}

		res = append(res, exView.ToDomain())
	}

	return res, nil
}

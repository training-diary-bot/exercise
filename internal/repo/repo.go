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
	"github.com/jackc/pgx/v5"
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

func (r Repo) CreateCustomExercise(ctx context.Context, userID uuid.UUID, name string) (uuid.UUID, error) {
	log := logger.New()

	req, args, err := sq.Insert(dbview.ExercisesTable).
		SetMap(map[string]any{
			dbview.ExercisesIDField:   exercise.ID,
			dbview.ExercisesNameField: exercise.Name,
		}).PlaceholderFormat(sq.Dollar).ToSql()
	// Добавить как разберусь
	//.Suffix(
	//	fmt.Sprintf("ON CONFLICT (%s) DO UPDATE SET ", dbview.ExercisesIDField) +
	//		fmt.Sprintf("RETURNING %s", dbview.ExercisesIDField),
	//)

	if err != nil {
		return uuid.UUID{}, errors.Wrap(err, "failed to create exercise")
	}
	log.Debug().Msgf("query: %s %q", req, args)

	var id uuid.UUID
	err = r.db.QueryRow(ctx, req, args...).Scan(&id)
	if err != nil {
		return uuid.UUID{}, errors.Wrapf(err, "%s %q", req, args)
	}

	return id, nil
}

func (r Repo) GetExercisesByMuscleGroup(ctx context.Context, userID uuid.UUID, MuscleGroupID uint64) ([]domain.Exercise, error) {
	log := logger.New()

	req, args, err := sq.
		Select(dbview.GetAllFields()...).
		From(dbview.ExercisesTable).
		Where(sq.Eq{
			dbview.ExercisesMuscleGroupField: mg,
		}).PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	log.Debug().Msgf("query: %s %q", req, args)

	exercises := make([]domain.Exercise, 0, 0)
	rows, err := r.db.Query(ctx, req, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query exec")
	}

	var exerciseView dbview.Exercise

	for rows.Next() {
		if err := pgxscan.ScanOne(&exerciseView, rows); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, domain.ErrExerciseNotFound
			}
			return nil, errors.Wrap(err, "failed to scan result")
		}

		if err != nil {
			return nil, errors.Wrap(err, "failed to map exercises to domain")
		}

		exercises = append(exercises, *exerciseView.ExerciseToDomain())
	}

	// посмотреть пресентер repo 87 строка. rows.Next()
	//exercises := make([]domain.Exercise)
	//
	//err := dbview.ExerciseToDomain(exerciseView)

	return exercises, nil
}

func (r Repo) GetCustomExercises(ctx context.Context, userID uuid.UUID) ([]domain.Exercise, error) {
	log := logger.New()

	req, args, err := sq.
		Select(dbview.GetAllFields()...).
		From(dbview.ExercisesTable).
		Where(sq.Eq{
			dbview.ExercisesMuscleGroupField: mg,
		}).PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	log.Debug().Msgf("query: %s %q", req, args)

	rows, err := r.db.Query(ctx, req, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query exec")
	}

	var exerciseView dbview.Exercise
	//тоже самое
	if err := pgxscan.ScanOne(&exerciseView, rows); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrExerciseNotFound
		}
		return nil, errors.Wrap(err, "failed to scan result")
	}

	exercise, err := dbview.ExerciseToDomain(exerciseView)
	if err != nil {
		return nil, errors.Wrap(err, "failed to map exercises to domain")
	}

	return exercise, nil
}

func (r Repo) GetMuscleGroups(ctx context.Context) ([]domain.MuscleGroup, error) {
	log := logger.New()

	req, args, err := sq.
		Select(dbview.GetAllFields()...).
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

	var muscleGroupsView dbview.MuscleGroups
	// тоже самое
	if err := pgxscan.ScanOne(&muscleGroupsView, rows); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrExerciseNotFound
		}
		return nil, errors.Wrap(err, "failed to scan result")
	}

	muscleGroups, err := dbview.MuscleGroupsToDomain(muscleGroupsView)
	if err != nil {
		return nil, errors.Wrap(err, "failed to map muscle groups to domain")
	}

	return muscleGroups, nil
}

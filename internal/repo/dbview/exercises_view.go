package dbview

const ExercisesTable = "exercises"
const MuscleGroupTable = "muscle_groups"

const (
	ExercisesIDField          = "id"
	ExercisesNameField        = "name"
	ExercisesMuscleGroupField = "muscle_group"
)

func GetAllFields() []string {
	return []string{
		ExercisesIDField,
		ExercisesNameField,
		ExercisesMuscleGroupField,
	}
}

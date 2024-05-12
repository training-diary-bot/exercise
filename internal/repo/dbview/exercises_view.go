package dbview

const ExercisesTable = "exercises"
const MuscleGroupTable = "muscle_groups"

const (
	ExercisesIDField          = "id"
	ExercisesNameField        = "name"
	ExercisesMuscleGroupField = "muscle_group_id"

	MuscleGroupsIDField   = "id"
	MuscleGroupsNameField = "name"
)

func GetExercisesAllFields() []string {
	return []string{
		ExercisesIDField,
		ExercisesNameField,
		ExercisesMuscleGroupField,
	}
}

func GetMuscleGroupsAllFields() []string {
	return []string{
		MuscleGroupsIDField,
		MuscleGroupsNameField,
	}
}

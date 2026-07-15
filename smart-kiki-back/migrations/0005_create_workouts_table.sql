-- +goose Up
CREATE TABLE workouts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    trainer_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    notes TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_workouts_trainer_id ON workouts (trainer_id);
CREATE INDEX idx_workouts_student_id ON workouts (student_id);

CREATE TABLE workout_exercises (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workout_id UUID NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
    exercise_id UUID NOT NULL REFERENCES exercises(id),
    position INT NOT NULL DEFAULT 0,
    sets INT NOT NULL,
    reps INT NOT NULL,
    load VARCHAR(50) NOT NULL DEFAULT '',
    rest_seconds INT NOT NULL DEFAULT 0,
    notes TEXT NOT NULL DEFAULT ''
);

CREATE INDEX idx_workout_exercises_workout_id ON workout_exercises (workout_id);

-- +goose Down
DROP TABLE IF EXISTS workout_exercises;
DROP TABLE IF EXISTS workouts;

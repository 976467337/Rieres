-- +goose Up
CREATE TABLE trainer_students (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    trainer_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (trainer_id, student_id)
);

CREATE INDEX idx_trainer_students_trainer_id ON trainer_students (trainer_id);

-- +goose Down
DROP TABLE IF EXISTS trainer_students;

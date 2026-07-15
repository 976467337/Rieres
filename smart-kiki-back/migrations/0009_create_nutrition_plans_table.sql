-- +goose Up
CREATE TABLE nutrition_plans (
    trainer_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL DEFAULT '',
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (trainer_id, student_id)
);

-- +goose Down
DROP TABLE IF EXISTS nutrition_plans;

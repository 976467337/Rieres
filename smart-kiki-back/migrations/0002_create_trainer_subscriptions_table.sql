-- +goose Up
CREATE TABLE trainer_subscriptions (
    trainer_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    plan VARCHAR(10) NOT NULL DEFAULT 'free' CHECK (plan IN ('free', 'plus', 'pro')),
    current_period_start TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS trainer_subscriptions;

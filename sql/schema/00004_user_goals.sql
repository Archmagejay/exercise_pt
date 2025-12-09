-- +goose Up
-- +goose StatementBegin
-- Create a new table 'user_goals' with a primary key and columns
CREATE TABLE user_goals (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE,
    goal_id UUID NOT NULL REFERENCES goals ON DELETE CASCADE,
    UNIQUE (user_id, goal_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop an existing table 'user_goals'
DROP TABLE IF EXISTS user_goals;
-- +goose StatementEnd

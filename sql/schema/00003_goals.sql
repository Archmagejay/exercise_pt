-- +goose Up
-- +goose StatementBegin
CREATE TABLE goals (
    id UUID PRIMARY KEY,
    type TEXT UNIQUE NOT NULL,
    goal_count INTEGER ARRAY[7],
    goal_dur TIME,
    goal_speed DECIMAL(4, 2),
    goal_tier INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE goals;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE goals (
    id UUID PRIMARY KEY,
    type TEXT UNIQUE NOT NULL,
    goal_plate_count INTEGER ARRAY[7],
    goal_dur TIME,
    goal_decimal DECIMAL(4, 2),
    goal_number INTEGER,
    goal_tier INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS goals;
-- +goose StatementEnd

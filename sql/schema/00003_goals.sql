-- +goose Up
-- +goose StatementBegin
CREATE TYPE goal_types AS ENUM (
    'Bench Press', --Plate count
    'Bike', -- Distance
    'Bisep Curls', --Plate count
    'Lateral Pulldown', --Plate count
    'Park Run', --Duration
    'Pectoral Fly', --Plate count
    'Plank', --Duration
    'Quad Curls', --Plate count
    'Treadmill',--Distance
    'Trapezius Lift', --Plate count
    'Trisep Curls', --Plate count
    'Waist',--Int
    'Weight'--Decimal
);

CREATE TABLE goals (
    id UUID PRIMARY KEY,
    goal_type goal_types NOT NULL,
    goal_plate_count INTEGER ARRAY[7],
    goal_dur INTERVAL NULL,
    goal_decimal DECIMAL(4, 2),
    goal_number INTEGER,
    goal_tier INTEGER NOT NULL,
    UNIQUE (goal_type, goal_tier)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS goals;
DROP TYPE IF EXISTS goal_types;
-- +goose StatementEnd

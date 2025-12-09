-- +goose Up
-- +goose StatementBegin
CREATE TABLE entrys (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE,
    date TIME NOT NULL,
    cardio DECIMAL(5, 2) NOT NULL,
    cardio_type BOOLEAN NOT NULL,
    plate_count INTEGER ARRAY[7] NOT NULL,
    plank_dur TIME NOT NULL,
    weight DECIMAL(6, 2) NOT NULL,
    waist DECIMAL (6, 2) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE entrys;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
ALTER TABLE entries ALTER COLUMN cardio_type DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE entries ALTER COLUMN cardio_type SET NOT NULL;
-- +goose StatementEnd

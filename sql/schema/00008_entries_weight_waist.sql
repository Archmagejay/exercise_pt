-- +goose Up
-- +goose StatementBegin
ALTER TABLE entries
    ALTER COLUMN weight DROP NOT NULL,
    ALTER COLUMN waist DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE entries
    ALTER COLUMN weight SET NOT NULL,
    ALTER COLUMN waist SET NOT NULL;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
ALTER TABLE entries
    ALTER COLUMN date TYPE TIMESTAMP WITH TIME ZONE
    USING
    TIMESTAMP WITH TIME ZONE 'epoch' + date;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE entries ALTER COLUMN date SET DATA TYPE TIME;
-- +goose StatementEnd

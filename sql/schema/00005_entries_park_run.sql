-- +goose Up
-- +goose StatementBegin
ALTER TABLE entries
ADD COLUMN park_run INTERVAL NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE entries
DROP COLUMN IF EXISTS park_run;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN access_token TEXT NOT NULL DEFAULT 'unset';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN access_token;
-- +goose StatementEnd

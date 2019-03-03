-- +goose Up
ALTER TABLE access_tokens
MODIFY access_token DROP DEFAULT;

ALTER TABLE access_tokens
MODIFY access_token NOT NULL;

-- +goose Down
ALTER TABLE access_tokens
MODIFY access_token NULL;

ALTER TABLE access_tokens
MODIFY access_token ADD DEFAULT NULL;

-- +goose Up
ALTER TABLE `access_tokens`
MODIFY `access_token` varchar(255) UNIQUE NOT NULL;

-- +goose Down
ALTER TABLE `access_tokens`
`access_token` varchar(255) DEFAULT NULL,

-- +goose Up
ALTER TABLE `users`
MODIFY `auth_token` varchar(255) NOT NULL;

ALTER TABLE `users`
DROP KEY `idx_users_auth_token`;

ALTER TABLE `users`
DROP KEY `auth_token`;

ALTER TABLE `users`
ADD INDEX `idx_auth_token`(`auth_token`);

-- +goose Down
ALTER TABLE `users`
MODIFY `auth_token` varchar(255) NOT NULL;

ALTER TABLE `users`
ADD KEY `idx_users_auth_token` (`auth_token`);

ALTER TABLE `users`
ADD UNIQUE KEY `auth_token`(`auth_token`);

ALTER TABLE `users`
DROP INDEX `idx_auth_token`;

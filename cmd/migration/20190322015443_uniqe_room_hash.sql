-- +goose Up
ALTER TABLE `chat_rooms`
MODIFY `room_hash` varchar(255) UNIQUE NOT NULL;

-- +goose Down
ALTER TABLE `access_tokens`
`room_hash` varchar(255) default NULL,

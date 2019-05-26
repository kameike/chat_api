-- +goose Up
DROP TABLE IF EXISTS `access_tokens`;
SET character_set_client = utf8mb4;
CREATE TABLE `access_tokens` (
  `user_id` int unsigned DEFAULT NULL,
  `access_token` varchar(255) DEFAULT NULL,
  UNIQUE KEY `access_token` (`access_token`),
  KEY `idx_access_tokens_user_id` (`user_id`),
  KEY `idx_access_tokens_access_token` (`access_token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET character_set_client = utf8mb4;
CREATE TABLE `chat_rooms` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `room_hash` varchar(255),
  `name` varchar(255), 
  PRIMARY KEY (`id`),
  KEY `idx_chat_rooms_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `messages`;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `messages` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `text` varchar(255),
  `user_id` int,
  `room_id` int,
  PRIMARY KEY (`id`),
  KEY `idx_messages_deleted_at` (`deleted_at`),
  KEY `idx_messages_user_id` (`user_id`),
  KEY `idx_messages_room_id` (`room_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `user_chat_rooms`;
SET character_set_client = utf8mb4;
CREATE TABLE `user_chat_rooms` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `user_id` int,
  `chat_room_id` int,
  `last_read_at` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_chat_rooms_chat_room_id` (`chat_room_id`),
  KEY `idx_user_chat_rooms_deleted_at` (`deleted_at`),
  KEY `idx_user_chat_rooms_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `users`;
SET character_set_client = utf8mb4;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `auth_token` varchar(255),
  `user_hash` varchar(255),
  `name` varchar(255) DEFAULT NULL,
  `url` text,
  `push_token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_token` (`auth_token`),
  UNIQUE KEY `user_hash` (`user_hash`),
  KEY `idx_users_auth_token` (`auth_token`),
  KEY `idx_users_user_hash` (`user_hash`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `access_tokens`;
DROP TABLE IF EXISTS `chat_rooms`;
DROP TABLE IF EXISTS `messages`;
DROP TABLE IF EXISTS `user_chat_rooms`;
DROP TABLE IF EXISTS `users`;

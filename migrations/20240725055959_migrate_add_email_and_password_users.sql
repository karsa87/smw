-- +goose Up
ALTER TABLE user ADD COLUMN email VARCHAR(255) NULL;
ALTER TABLE user ADD COLUMN password TEXT NULL;
ALTER TABLE `user` ADD INDEX `email_index` (`email`);

-- +goose Down
DROP INDEX `email_index` ON `user`;
ALTER TABLE user DROP COLUMN email;
ALTER TABLE user DROP COLUMN password;

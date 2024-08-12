-- +goose Up
CREATE TABLE IF NOT EXISTS `user`(
    `id` int UNIQUE PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `address` text NULL,
    `gender` ENUM('male', 'female') DEFAULT 'male'
) ENGINE = InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;


-- +goose Down
DROP TABLE IF EXISTS `user`;
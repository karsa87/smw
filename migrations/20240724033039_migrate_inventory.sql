-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `inventory`(
    `id` int UNIQUE PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `name` varchar(255) NOT NULL,
    `stock` int UNSIGNED DEFAULT 0,
    `price` decimal(18, 2) UNSIGNED DEFAULT 0,
    `description` text NULL,
    FOREIGN KEY(`user_id`) 
        REFERENCES user(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FULLTEXT INDEX (`name`)
) ENGINE = InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `inventory`;
-- +goose StatementEnd

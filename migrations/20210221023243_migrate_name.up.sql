-- +goose Up
CREATE TABLE IF NOT EXISTS `history`(
    `id` int UNIQUE PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `source` varchar(255),
    `destination` varchar(255),
    `original` varchar(255),
    `translation` varchar(255)
) ENGINE = InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;

-- +goose Down
DROP TABLE IF EXISTS `history`;
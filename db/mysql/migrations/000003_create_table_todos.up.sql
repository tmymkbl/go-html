DROP TABLE IF EXISTS `todos`;
CREATE TABLE IF NOT EXISTS `todos`(
    `id` INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `title` VARCHAR(64),
    `content` TEXT,
    `user_id` INTEGER UNSIGNED,
    `completed_at` TIMESTAMP,
    `created_at`   TIMESTAMP,
    `updated_at`   TIMESTAMP,
    `deleted_at`   TIMESTAMP);

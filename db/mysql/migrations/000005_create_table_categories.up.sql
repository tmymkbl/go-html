DROP TABLE IF EXISTS `categories`;
CREATE TABLE IF NOT EXISTS `categories`(
    `id` INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `categories` VARCHAR(64),
    `created_at` TIMESTAMP,
    `updated_at` TIMESTAMP,
    `deleted_at` TIMESTAMP);

DROP TABLE IF EXISTS `categories_blogs`;
CREATE TABLE IF NOT EXISTS `categories_blogs`(
    `id` INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `categories_id` INTEGER UNSIGNED,
    `blog_id` INTEGER UNSIGNED,
    `created_at` TIMESTAMP,
    `updated_at` TIMESTAMP,
    `deleted_at` TIMESTAMP);

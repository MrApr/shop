CREATE TABLE `user_addresses`
(
    `id`         INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `user_id`    INTEGER UNSIGNED                NOT NULL,
    `city_id`    INTEGER UNSIGNED                NOT NULL,
    `address`    VARCHAR(500)                    NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`city_id`) REFERENCES `cities` (`id`)
);
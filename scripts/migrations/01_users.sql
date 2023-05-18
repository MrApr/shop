CREATE TABLE `users`
(
    `id`            INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `phone_number`  VARCHAR(255)                    NOT NULL,
    `name`          VARCHAR(255),
    `uuid`          VARCHAR(255)                    NOT NULL,
    `password`      VARCHAR(255)                    NOT NULL,
    `last_login_at` TIMESTAMP,
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE UNIQUE INDEX users_phone_unique_index ON `users` (`phone_number`);
CREATE UNIQUE INDEX users_uuid_unique_index ON `users` (`uuid`);
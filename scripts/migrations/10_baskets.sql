CREATE TABLE `baskets`
(
    `id`         INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `user_id`    INTEGER UNSIGNED                NOT NULL,
    `status`     BOOL      DEFAULT TRUE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at`  TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE INDEX basket_status_index ON `baskets` (`status`);
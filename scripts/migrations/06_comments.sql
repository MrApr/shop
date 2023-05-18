CREATE TABLE `comments`
(
    `id`          INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `product_id`  INTEGER UNSIGNED                NOT NULL,
    `user_id`     INTEGER UNSIGNED                NOT NULL,
    `description` VARCHAR(255)                    NOT NULL,
    `status`      BOOL      DEFAULT FALSE,
    `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE INDEX comments_status_index ON `comments` (`status`);
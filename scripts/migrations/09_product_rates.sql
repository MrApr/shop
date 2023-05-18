CREATE TABLE `product_rates`
(
    `id`         INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `product_id` INTEGER UNSIGNED                NOT NULL,
    `user_id`    INTEGER UNSIGNED                NOT NULL,
    `score`      SMALLINT UNSIGNED               NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE INDEX product_rates_score_index ON `product_rates` (`score`);
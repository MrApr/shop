CREATE TABLE `likeables`
(
    `product_id` INTEGER UNSIGNED NOT NULL,
    `user_id`    INTEGER UNSIGNED NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);

CREATE UNIQUE INDEX likeables_unique_index ON `likeables` (`product_id`, `user_id`);
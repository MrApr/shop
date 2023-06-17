CREATE TABLE `dislikes`
(
    `product_id` INTEGER UNSIGNED NOT NULL,
    `user_id`    INTEGER UNSIGNED NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);

CREATE UNIQUE INDEX dislikeables_unique_index ON `dislikes` (`product_id`, `user_id`);
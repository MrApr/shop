CREATE TABLE `user_alerts`
(
    `user_id`    INTEGER UNSIGNED NOT NULL,
    `product_id` INTEGER UNSIGNED NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);
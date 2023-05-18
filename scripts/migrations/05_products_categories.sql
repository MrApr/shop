CREATE TABLE `product_categories_table`
(
    `product_id`  INTEGER UNSIGNED NOT NULL,
    `category_id` INTEGER UNSIGNED NOT NULL,
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
);
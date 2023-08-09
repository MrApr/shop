CREATE TABLE `product_categories`
(
    `product_id`  INTEGER UNSIGNED NOT NULL,
    `category_id` INTEGER UNSIGNED NOT NULL,
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
);
CREATE UNIQUE INDEX product_cat_unique ON `product_categories`(`product_id`, category_id);
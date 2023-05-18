CREATE TABLE `basket_products`
(
    `basket_id`  INTEGER UNSIGNED   NOT NULL,
    `product_id` INTEGER UNSIGNED   NOT NULL,
    `amount`     MEDIUMINT UNSIGNED NOT NULL,
    `unit_price` DOUBLE UNSIGNED    NOT NULL,
    FOREIGN KEY (`basket_id`) REFERENCES `baskets` (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);
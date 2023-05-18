CREATE TABLE `views`
(
    `id`         INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `product_id` INTEGER UNSIGNED                NOT NULL,
    `IP`         VARCHAR(255)                    NOT NULL,
    `agent`      VARCHAR(255),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);
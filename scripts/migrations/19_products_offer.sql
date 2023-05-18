CREATE TABLE `products_offer`
(
    `id`          INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `type_id`     INTEGER UNSIGNED                NOT NULL,
    `category_id` INTEGER UNSIGNED                NOT NULL,
    `title`       VARCHAR(255)                    NOT NULL,
    `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
    FOREIGN KEY (`type_id`) REFERENCES `types` (`id`)
);
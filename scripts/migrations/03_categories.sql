CREATE TABLE `categories`
(
    `id`            INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `type_id`       INTEGER UNSIGNED,
    `parent_cat_id` INTEGER UNSIGNED,
    `title`         VARCHAR(255)                    NOT NULL,
    `indent`        MEDIUMINT UNSIGNED              NOT NULL,
    `order`         MEDIUMINT UNSIGNED              NOT NULL,
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`parent_cat_id`) REFERENCES `categories` (`id`),
    FOREIGN KEY (`type_id`) REFERENCES `types` (`id`)
);

CREATE INDEX categories_indent_order_index ON `categories` (`indent`, `order`);
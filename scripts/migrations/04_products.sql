CREATE TABLE `products`
(
    `id`          INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `title`       VARCHAR(255)                    NOT NULL,
    `code`        VARCHAR(255)                    NOT NULL,
    `amount`      MEDIUMINT UNSIGNED              NOT NULL,
    `price`       DOUBLE UNSIGNED                 NOT NULL,
    `weight`      INTEGER UNSIGNED,
    `description` Text,
    `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`   TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE UNIQUE INDEX products_code_unique_index ON `products` (`code`);
CREATE INDEX products_title_index ON `products` (`title`);
CREATE INDEX products_price_index ON `products` (`price`);
CREATE FULLTEXT INDEX products_description_index ON `products` (`description`);
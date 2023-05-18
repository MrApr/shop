CREATE TABLE `discount_codes`
(
    `id`               INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `title`            VARCHAR(255)                    NOT NULL,
    `code`             VARCHAR(255)                    NOT NULL,
    `discount_percent` INTEGER UNSIGNED                NOT NULL,
    `status`           BOOL      DEFAULT TRUE,
    `created_at`       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at`        TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE UNIQUE INDEX discount_codes_codes_unique ON `discount_codes` (`code`);
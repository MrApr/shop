CREATE TABLE `gateway_types`
(
    `id`         INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `title`      VARCHAR(255)                    NOT NULL,
    `status`     BOOL      DEFAULT TRUE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE INDEX gateway_types_status_index ON `gateway_types` (`status`);
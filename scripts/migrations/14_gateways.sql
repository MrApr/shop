CREATE TABLE `gateways`
(
    `id`              INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `name`            VARCHAR(255)                    NOT NULL,
    `gateway_type_id` INTEGER UNSIGNED                NOT NULL,
    `token`           VARCHAR(255)                    NOT NULL,
    `status`          BOOL      DEFAULT FALSE,
    `created_at`      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at`       TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`gateway_type_id`) REFERENCES `gateway_types` (`id`)
);

CREATE INDEX gateways_status_index ON `gateways` (`status`);
CREATE TABLE `post_types`
(
    `id`               INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `title`            VARCHAR(255)                    NOT NULL,
    `price`            DOUBLE UNSIGNED                 NOT NULL,
    `deliverable_time` TIMESTAMP                       NOT NULL,
    `created_at`       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at`        TIMESTAMP,
    PRIMARY KEY (`id`)
);
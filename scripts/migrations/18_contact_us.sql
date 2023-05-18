CREATE TABLE `contact_us`
(
    `id`          INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `email`       VARCHAR(255)                    NOT NULL,
    `title`       VARCHAR(255)                    NOT NULL,
    `description` VARCHAR(500)                    NOT NULL,
    `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
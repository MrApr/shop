CREATE TABLE `dislikeables`
(
    `dislikeable_id`   INTEGER UNSIGNED NOT NULL,
    `dislikeable_type` VARCHAR(255)     NOT NULL,
    `user_id`       INTEGER UNSIGNED NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE INDEX dislikeables_index ON `dislikeables` (`dislikeable_id`, `dislikeable_type`);
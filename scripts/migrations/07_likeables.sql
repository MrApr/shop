CREATE TABLE `likeables`
(
    `likeable_id`   INTEGER UNSIGNED NOT NULL,
    `likeable_type` VARCHAR(255)     NOT NULL,
    `user_id`       INTEGER UNSIGNED NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE INDEX likeables_index ON `likeables` (`likeable_id`, `likeable_type`);
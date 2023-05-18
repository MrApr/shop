CREATE TABLE `payments`
(
    `id`           INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `user_id`      INTEGER UNSIGNED                NOT NULL,
    `basket_id`    INTEGER UNSIGNED                NOT NULL,
    `address_id`   INTEGER UNSIGNED,
    `discount_id`  INTEGER UNSIGNED,
    `gateway_id`   INTEGER UNSIGNED,
    `post_type_id` INTEGER UNSIGNED,
    `total_price`  DOUBLE UNSIGNED                 NOT NULL,
    `ref_num`      VARCHAR(255),
    `trace_num`    VARCHAR(255),
    `status`       BOOL                            NOT NULL DEFAULT FALSE,
    `created_at`   TIMESTAMP                                DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   TIMESTAMP                                DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`basket_id`) REFERENCES `baskets` (`id`),
    FOREIGN KEY (`discount_id`) REFERENCES `discount_codes` (`id`),
    FOREIGN KEY (`address_id`) REFERENCES `user_addresses` (`id`),
    FOREIGN KEY (`gateway_id`) REFERENCES `gateways` (`id`),
    FOREIGN KEY (`post_type_id`) REFERENCES `post_types` (`id`)
);

CREATE UNIQUE INDEX payments_ref_num_unique ON `payments` (`ref_num`);
CREATE UNIQUE INDEX payments_trace_num_unique ON `payments` (`trace_num`);
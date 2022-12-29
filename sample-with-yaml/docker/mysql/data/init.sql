DROP TABLE IF EXISTS `users`;
create table `users` (
    `user_id`         BIGINT(20) AUTO_INCREMENT,
    `user_name`       VARCHAR(36) NOT NULL,
    `user_created_at` DATETIME DEFAULT '0000-00-00 00:00:00',
    UNIQUE KEY uq_keys(user_name),
    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO users (user_name, user_created_at) VALUES ('Michael', now());

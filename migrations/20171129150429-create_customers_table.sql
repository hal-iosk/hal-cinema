
-- +migrate Up
CREATE TABLE `customers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `first_name_read` varchar(255) NOT NULL,
  `last_name_read` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `birthdate` varchar(255) NOT NULL,
  `magazine` tinyint(1) NOT NULL,
  `point_count` bigint(20) NOT NULL,
  `credit_card_name` varchar(255) NOT NULL,
  `credit_card_limit` varchar(255) NOT NULL,
  `credit_card_number` varchar(255) NOT NULL,
  `security_code` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `customers`;

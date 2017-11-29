
-- +migrate Up
CREATE TABLE `prices` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `customer_type` varchar(255) NOT NULL,
  `price` bigint(20) NOT NULL,
  `last_updated_administrator_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_prices_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `prices`;

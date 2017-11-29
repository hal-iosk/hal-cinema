
-- +migrate Up
CREATE TABLE `earnings` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `reserve_id` varchar(255) NOT NULL,
  `earning_date` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `reserve_id` (`reserve_id`),
  KEY `idx_earnings_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `earnings`;

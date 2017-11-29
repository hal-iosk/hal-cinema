
-- +migrate Up
CREATE TABLE `reserves` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `reserve_number` bigint(20) NOT NULL,
  `movie_name` varchar(255) NOT NULL,
  `reserve_time` varchar(255) NOT NULL,
  `watch_date` varchar(255) NOT NULL,
  `watch_start_time` varchar(255) NOT NULL,
  `ticket_time` varchar(255) NOT NULL,
  `theater_number` bigint(20) NOT NULL,
  `price_id` varchar(255) NOT NULL,
  `seat_number` varchar(255) NOT NULL,
  `customer_id` varchar(255) NOT NULL,
  `cancel_check` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `price_id` (`price_id`),
  UNIQUE KEY `customer_id` (`customer_id`),
  KEY `idx_reserves_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `reserves`;

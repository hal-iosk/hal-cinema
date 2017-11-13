-- +migrate Up
CREATE TABLE `administrators` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `admin_id` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `authority_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_id` (`admin_id`),
  UNIQUE KEY `authority_id` (`authority_id`),
  KEY `idx_administrators_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `authorities` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `authority_id` varchar(255) NOT NULL,
  `authority_name` varchar(255) NOT NULL,
  `authority_code` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `authority_id` (`authority_id`),
  KEY `idx_authorities_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `customers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `customer_id` varchar(255) NOT NULL,
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
  UNIQUE KEY `customer_id` (`customer_id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `earnings` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `earning_id` varchar(255) NOT NULL,
  `reserve_id` varchar(255) NOT NULL,
  `earning_date` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `earning_id` (`earning_id`),
  UNIQUE KEY `reserve_id` (`reserve_id`),
  KEY `idx_earnings_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `movies` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `movie_id` varchar(255) NOT NULL,
  `movie_name` varchar(255) NOT NULL,
  `start_date` varchar(255) NOT NULL,
  `end_date` varchar(255) NOT NULL,
  `watch_time` bigint(20) NOT NULL,
  `administrator_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `movie_id` (`movie_id`),
  KEY `idx_movies_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `prices` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `price_id` varchar(255) NOT NULL,
  `customer_type` varchar(255) NOT NULL,
  `price` bigint(20) NOT NULL,
  `last_updated_administrator_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `price_id` (`price_id`),
  KEY `idx_prices_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `reserves` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `reserve_id` varchar(255) NOT NULL,
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
  UNIQUE KEY `reserve_id` (`reserve_id`),
  UNIQUE KEY `price_id` (`price_id`),
  UNIQUE KEY `customer_id` (`customer_id`),
  KEY `idx_reserves_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `screening_schedules` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `schedule_id` varchar(255) NOT NULL,
  `movie_id` varchar(255) NOT NULL,
  `watch_start_time` varchar(255) NOT NULL,
  `watch_date` varchar(255) NOT NULL,
  `theater_number` bigint(20) NOT NULL,
  `last_updated_administrator_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `schedule_id` (`schedule_id`),
  UNIQUE KEY `movie_id` (`movie_id`),
  KEY `idx_screening_schedules_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `administrators`;
DROP TABLE `authorities`;
DROP TABLE `customers`;
DROP TABLE `earnings`;
DROP TABLE `movies`;
DROP TABLE `prices`;
DROP TABLE `reserves`;
DROP TABLE `screening_schedules`;

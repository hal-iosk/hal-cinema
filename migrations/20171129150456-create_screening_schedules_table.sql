
-- +migrate Up
CREATE TABLE `screening_schedules` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `movie_id` varchar(255) NOT NULL,
  `watch_start_time` varchar(255) NOT NULL,
  `watch_date` varchar(255) NOT NULL,
  `theater_number` bigint(20) NOT NULL,
  `last_updated_administrator_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `movie_id` (`movie_id`),
  KEY `idx_screening_schedules_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `screening_schedules`;

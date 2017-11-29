
-- +migrate Up
CREATE TABLE `movies` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `movie_name` varchar(255) NOT NULL,
  `start_date` varchar(255) NOT NULL,
  `end_date` varchar(255) NOT NULL,
  `watch_time` bigint(20) NOT NULL,
  `administrator_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_movies_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `movies`;

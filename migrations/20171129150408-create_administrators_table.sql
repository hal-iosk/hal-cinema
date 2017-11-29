
-- +migrate Up
CREATE TABLE `administrators` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `authority_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `authority_id` (`authority_id`),
  KEY `idx_administrators_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `administrators`;

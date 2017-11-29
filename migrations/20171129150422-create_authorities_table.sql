
-- +migrate Up
CREATE TABLE `authorities` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `authority_name` varchar(255) NOT NULL,
  `authority_code` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_authorities_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `authorities`;

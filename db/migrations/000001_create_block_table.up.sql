CREATE TABLE IF NOT EXISTS `Block` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `number` bigint(20) NOT NULL,
  `hash` varchar(100) NOT NULL DEFAULT '',
  `timestamp` bigint(20) NOT NULL DEFAULT '0',
  `parentHash` varchar(100) NOT NULL DEFAULT '',
  `isRemoved` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uni_hash` (`hash` ASC),
  INDEX `idx_number_isRemoved` (`number`, `isRemoved`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

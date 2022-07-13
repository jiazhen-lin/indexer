CREATE TABLE IF NOT EXISTS `TransactionLog` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `txHash` varchar(100) NOT NULL DEFAULT '',
  `idx` int(10) NOT NULL DEFAULT '0',
  `data` text,
  PRIMARY KEY (`id`),
  INDEX `idx_txHash_idx` (`txHash`, `idx`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

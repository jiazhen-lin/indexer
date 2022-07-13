CREATE TABLE IF NOT EXISTS `Transaction` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `txHash` varchar(100) NOT NULL DEFAULT '',
  `blockHash` varchar(100) NOT NULL DEFAULT '',
  `idx` int(10) NOT NULL DEFAULT '0',
  `from` varchar(100) NOT NULL DEFAULT '',
  `to` varchar(100) NOT NULL DEFAULT '',
  `nonce` int(10) NOT NULL DEFAULT '0',
  `value` varchar(100) NOT NULL DEFAULT '',
  `data` text,
  PRIMARY KEY (`id`),
  INDEX `idx_txHash` (`txHash`),
  INDEX `idx_blockHash_idx` (`blockHash`, `idx`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

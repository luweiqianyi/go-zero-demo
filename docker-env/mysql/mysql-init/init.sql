CREATE TABLE `TbUserAccount` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                              `accountName` varchar(191) DEFAULT NULL,
                              `password` varchar(191) DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `accountName` (`accountName`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8
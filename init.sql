-- MariaDB dump 10.19  Distrib 10.6.4-MariaDB, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: kyoto
-- ------------------------------------------------------
-- Server version	5.7.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `birthdays`
--

DROP TABLE IF EXISTS `birthdays`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `birthdays`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `animate`    longtext COLLATE utf8mb4_unicode_ci,
    `name`       longtext COLLATE utf8mb4_unicode_ci,
    `date`       longtext COLLATE utf8mb4_unicode_ci,
    PRIMARY KEY (`id`),
    KEY          `idx_birthdays_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `birthdays`
--

LOCK
TABLES `birthdays` WRITE;
/*!40000 ALTER TABLE `birthdays` DISABLE KEYS */;
INSERT INTO `birthdays`
VALUES (27, '2021-08-22 19:27:58.963', '2021-08-22 19:27:58.963', NULL, '吹响吧！上低音号', '久石奏', '01-07'),
       (28, '2021-08-22 19:27:58.981', '2021-08-22 19:27:58.981', NULL, '轻音少女', '秋山澪', '01-15'),
       (29, '2021-08-22 19:27:58.987', '2021-08-22 19:27:58.987', NULL, '吹响吧！上低音号', '黄前麻美子', '01-25'),
       (30, '2021-08-22 19:27:58.995', '2021-08-22 19:27:58.995', NULL, '轻音少女', '山中佐和子', '01-31'),
       (31, '2021-08-22 19:27:59.000', '2021-08-22 19:27:59.000', NULL, '吹响吧！上低音号', '加藤叶月', '02-13'),
       (32, '2021-08-22 19:27:59.006', '2021-08-22 19:27:59.006', NULL, '日常', '水上麻衣', '02-17'),
       (33, '2021-08-22 19:27:59.014', '2021-08-22 19:27:59.014', NULL, '轻音少女', '平泽忧', '02-22'),
       (34, '2021-08-22 19:27:59.021', '2021-08-22 19:27:59.021', NULL, '吹响吧！上低音号', '吉川优子', '04-15'),
       (35, '2021-08-22 19:27:59.030', '2021-08-22 19:27:59.030', NULL, '吹响吧！上低音号', '高坂丽奈', '05-15'),
       (36, '2021-08-22 19:27:59.036', '2021-08-22 19:27:59.036', NULL, '声之形', '西宫硝子', '06-07'),
       (37, '2021-08-22 19:27:59.044', '2021-08-22 19:27:59.044', NULL, '吹响吧！上低音号', '中川夏纪', '06.23'),
       (38, '2021-08-22 19:27:59.052', '2021-08-22 19:27:59.052', NULL, '轻音少女', '琴吹紬', '07-02'),
       (39, '2021-08-22 19:27:59.061', '2021-08-22 19:27:59.061', NULL, '莉兹与青鸟', '铠冢霙', '07-02'),
       (40, '2021-08-22 19:27:59.070', '2021-08-22 19:27:59.070', NULL, '吹响吧！上低音号', '黄前久美子', '08-21'),
       (41, '2021-08-22 19:27:59.080', '2021-08-22 19:27:59.080', NULL, '轻音少女', '田井中律', '08-21'),
       (42, '2021-08-22 19:27:59.085', '2021-08-22 19:27:59.085', NULL, '吹响吧！上低音号', '泷升', '08-23'),
       (43, '2021-08-22 19:27:59.091', '2021-08-22 19:27:59.091', NULL, '吹响吧！上低音号', '中世古香织', '09-03'),
       (44, '2021-08-22 19:27:59.100', '2021-08-22 19:27:59.100', NULL, '吹响吧！上低音号', '冢本秀一', '09-18'),
       (45, '2021-08-22 19:27:59.105', '2021-08-22 19:27:59.105', NULL, '吹响吧！上低音号', '小笠原晴香', '10-28'),
       (46, '2021-08-22 19:27:59.115', '2021-08-22 19:27:59.115', NULL, '吹响吧！上低音号', '川岛绿辉', '11-03'),
       (47, '2021-08-22 19:27:59.125', '2021-08-22 19:27:59.125', NULL, '日常', '东云名乃', '11-07'),
       (48, '2021-08-22 19:27:59.136', '2021-08-22 19:27:59.136', NULL, '轻音少女', '中野梓', '11-11'),
       (49, '2021-08-22 19:27:59.147', '2021-08-22 19:27:59.147', NULL, '轻音少女', '平泽唯', '11-27'),
       (50, '2021-08-22 19:27:59.156', '2021-08-22 19:27:59.156', NULL, '莉兹与青鸟', '伞木希美', '12-03'),
       (51, '2021-08-22 19:27:59.165', '2021-08-22 19:27:59.165', NULL, '吹响吧！上低音号', '田中明日香', '12-25'),
       (52, '2021-08-22 19:27:59.170', '2021-08-22 19:27:59.170', NULL, '轻音少女', '真锅和', '12-26');
/*!40000 ALTER TABLE `birthdays` ENABLE KEYS */;
UNLOCK
TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-08-22 19:35:37

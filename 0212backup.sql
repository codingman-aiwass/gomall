-- MySQL dump 10.13  Distrib 8.4.3, for Linux (aarch64)
--
-- Host: localhost    Database: go_mall
-- ------------------------------------------------------
-- Server version	8.4.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `address_models`
--

DROP TABLE IF EXISTS `address_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `address_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_at` datetime(3) DEFAULT NULL,
  `user_id` int unsigned NOT NULL,
  `street_address` varchar(32) NOT NULL,
  `city` varchar(32) NOT NULL,
  `state` varchar(32) DEFAULT NULL,
  `country` varchar(32) NOT NULL,
  `zipcode` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_address_models_delete_at` (`delete_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `address_models`
--

LOCK TABLES `address_models` WRITE;
/*!40000 ALTER TABLE `address_models` DISABLE KEYS */;
INSERT INTO `address_models` VALUES (1,'2025-01-13 07:17:34',NULL,2,'No.2 Daxue Road','zhuhai','guangdong','china',582011),(2,'2025-01-14 14:27:41',NULL,2,'No.3 Daxue Road','zhuhai','guangdong','china',582021);
/*!40000 ALTER TABLE `address_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cart_models`
--

DROP TABLE IF EXISTS `cart_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cart_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_at` datetime(3) DEFAULT NULL,
  `user_id` int unsigned NOT NULL,
  `product_id` int unsigned NOT NULL,
  `quantity` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_cart_models_delete_at` (`delete_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cart_models`
--

LOCK TABLES `cart_models` WRITE;
/*!40000 ALTER TABLE `cart_models` DISABLE KEYS */;
INSERT INTO `cart_models` VALUES (1,'2025-01-10 06:57:55',NULL,1,1,2),(2,'2025-01-10 07:33:59','2025-01-10 15:44:20.240',2,1,4),(3,'2025-01-10 07:35:52','2025-01-10 15:49:34.389',2,2,8),(4,'2025-01-10 07:53:47',NULL,2,2,1);
/*!40000 ALTER TABLE `cart_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (4,'g','test1@qq.com','user',NULL,NULL,NULL,NULL),(21,'g','test2@qq.com','merchant',NULL,NULL,NULL,NULL),(1,'p','admin','/admin','GET',NULL,NULL,NULL),(5,'p','admin','/admin','POST',NULL,NULL,NULL),(18,'p','merchant','/api/product/create','POST',NULL,NULL,NULL),(20,'p','merchant','/api/product/remove','POST',NULL,NULL,NULL),(19,'p','merchant','/api/product/update','POST',NULL,NULL,NULL),(2,'p','merchant','/merchant','GET',NULL,NULL,NULL),(11,'p','user','/api/cart','GET',NULL,NULL,NULL),(17,'p','user','/api/cart','POST',NULL,NULL,NULL),(9,'p','user','/api/order','GET',NULL,NULL,NULL),(15,'p','user','/api/order','POST',NULL,NULL,NULL),(8,'p','user','/api/payment','GET',NULL,NULL,NULL),(14,'p','user','/api/payment','POST',NULL,NULL,NULL),(12,'p','user','/api/product/detail','GET',NULL,NULL,NULL),(7,'p','user','/api/user','GET',NULL,NULL,NULL),(13,'p','user','/api/user','POST',NULL,NULL,NULL),(3,'p','user','/user','GET',NULL,NULL,NULL),(6,'p','user','/user','POST',NULL,NULL,NULL);
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category_models`
--

DROP TABLE IF EXISTS `category_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_at` datetime(3) DEFAULT NULL,
  `product_id` int unsigned NOT NULL,
  `category` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_category_models_delete_at` (`delete_at`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category_models`
--

LOCK TABLES `category_models` WRITE;
/*!40000 ALTER TABLE `category_models` DISABLE KEYS */;
INSERT INTO `category_models` VALUES (1,'2025-01-09 13:45:45',NULL,1,'c1'),(2,'2025-01-09 13:45:45',NULL,1,'c2'),(3,'2025-01-09 13:46:52',NULL,2,'c1'),(4,'2025-01-09 13:46:52',NULL,2,'c2'),(5,'2025-01-09 13:46:59',NULL,3,'c1'),(6,'2025-01-09 13:46:59',NULL,3,'c2'),(7,'2025-01-09 13:47:16',NULL,4,'c3'),(8,'2025-01-09 13:47:30','2025-01-09 22:02:21.473',5,'c3'),(9,'2025-01-09 13:47:30','2025-01-09 22:02:21.473',5,'c4'),(10,'2025-01-09 13:49:53',NULL,4,'c3'),(11,'2025-01-09 13:49:53',NULL,4,'c4'),(12,'2025-01-15 16:13:30',NULL,6,'c3'),(13,'2025-01-15 16:13:30',NULL,6,'c4'),(14,'2025-01-15 16:15:01',NULL,7,'c3'),(15,'2025-01-15 16:15:01',NULL,7,'c6');
/*!40000 ALTER TABLE `category_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_item_models`
--

DROP TABLE IF EXISTS `order_item_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_item_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_at` datetime(3) DEFAULT NULL,
  `order_id` int unsigned NOT NULL,
  `product_id` int unsigned NOT NULL,
  `quantity` int unsigned NOT NULL,
  `cost` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_order_item_models_delete_at` (`delete_at`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_item_models`
--

LOCK TABLES `order_item_models` WRITE;
/*!40000 ALTER TABLE `order_item_models` DISABLE KEYS */;
INSERT INTO `order_item_models` VALUES (1,'2025-01-13 07:17:34',NULL,1,2,2,123),(2,'2025-01-13 07:52:29',NULL,2,3,3,129),(7,'2025-01-13 12:35:00',NULL,7,1,3,129),(8,'2025-01-13 12:44:05',NULL,8,3,3,129),(9,'2025-01-13 12:49:30',NULL,9,6,3,129),(10,'2025-01-13 13:15:25',NULL,10,66,3,129),(11,'2025-01-14 14:27:41',NULL,11,3,3,12123),(12,'2025-01-14 14:30:54',NULL,12,3,3,12123),(13,'2025-01-14 14:32:12',NULL,13,3,3,12123),(14,'2025-01-14 14:34:55',NULL,14,3,3,12123),(15,'2025-01-14 14:35:58',NULL,15,3,3,12123),(16,'2025-01-14 14:42:36',NULL,16,3,3,12123),(17,'2025-01-14 14:43:50',NULL,17,3,3,12123),(18,'2025-01-14 14:49:56',NULL,18,3,3,12123),(19,'2025-01-14 14:51:42',NULL,19,3,3,12123),(20,'2025-01-14 14:52:56',NULL,20,3,3,12123),(21,'2025-01-14 14:54:50',NULL,21,3,3,12123),(23,'2025-01-15 06:49:41',NULL,23,3,3,12123),(24,'2025-01-15 06:59:30',NULL,24,3,3,12123),(25,'2025-01-15 07:08:21',NULL,25,3,3,12123),(26,'2025-01-15 07:18:39',NULL,26,3,3,12123),(27,'2025-01-15 07:20:26',NULL,27,3,3,12123),(28,'2025-01-15 07:23:38',NULL,28,3,3,12123),(29,'2025-01-15 07:24:14',NULL,29,3,3,12123),(30,'2025-01-15 07:24:30',NULL,30,3,3,12123),(31,'2025-01-15 07:35:07',NULL,31,3,3,12123),(32,'2025-01-15 07:36:09',NULL,32,3,3,12123),(33,'2025-01-15 07:36:13',NULL,33,3,3,12123),(34,'2025-01-15 07:39:21',NULL,34,3,3,12123),(35,'2025-01-15 07:39:49',NULL,35,3,3,12123),(36,'2025-01-15 07:39:54',NULL,36,3,3,12123),(37,'2025-01-15 07:41:19',NULL,37,3,3,12123),(38,'2025-01-15 07:43:12',NULL,38,3,3,12123),(39,'2025-01-15 07:43:34',NULL,39,3,3,12123),(40,'2025-01-15 07:43:42',NULL,40,3,3,12123),(41,'2025-01-15 07:43:47',NULL,41,3,3,12123),(42,'2025-01-15 07:59:43',NULL,42,3,3,12123),(43,'2025-01-15 07:59:57',NULL,43,3,3,12123),(44,'2025-01-15 08:03:13',NULL,44,3,3,12123),(45,'2025-01-15 08:04:34',NULL,45,3,3,12123),(46,'2025-01-15 08:04:48',NULL,46,3,3,12123),(47,'2025-01-15 08:09:47',NULL,47,3,3,12123),(48,'2025-01-15 08:09:56',NULL,48,3,3,12123),(49,'2025-01-15 08:10:01',NULL,49,3,3,12123),(50,'2025-01-15 08:10:33',NULL,50,1,3,12123),(51,'2025-01-15 08:10:42',NULL,51,1,3,12123),(52,'2025-01-15 08:13:44',NULL,52,1,3,12123),(53,'2025-01-15 08:14:35',NULL,53,1,3,12123),(54,'2025-01-15 08:15:27',NULL,54,1,3,12123),(55,'2025-01-15 08:15:53',NULL,55,1,3,12123),(56,'2025-01-15 08:16:18',NULL,56,1,3,12123),(57,'2025-01-15 08:16:27',NULL,57,1,3,12123),(58,'2025-01-15 08:16:32',NULL,58,1,3,12123),(59,'2025-01-15 08:17:21',NULL,59,1,3,12123),(60,'2025-01-15 08:19:20',NULL,60,1,3,12123),(61,'2025-01-15 08:19:29',NULL,61,1,3,12123),(62,'2025-01-15 08:21:22',NULL,62,1,3,12123),(63,'2025-01-15 08:21:30',NULL,63,1,3,12123),(64,'2025-01-15 09:04:15',NULL,64,1,3,12123),(65,'2025-01-15 09:08:25',NULL,65,1,3,12123),(66,'2025-01-15 09:08:45',NULL,66,1,3,12123),(67,'2025-01-15 09:08:50',NULL,67,1,3,12123),(68,'2025-01-15 09:09:08',NULL,68,1,3,12123),(69,'2025-01-15 09:09:14',NULL,69,1,3,12123),(70,'2025-01-15 09:09:18',NULL,70,1,3,12123),(71,'2025-01-15 09:09:22',NULL,71,1,3,12123),(72,'2025-01-15 09:09:30',NULL,72,1,3,12123);
/*!40000 ALTER TABLE `order_item_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_models`
--

DROP TABLE IF EXISTS `order_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_at` datetime(3) DEFAULT NULL,
  `user_id` int unsigned NOT NULL,
  `user_currency` varchar(32) NOT NULL,
  `address_id` int unsigned NOT NULL,
  `email` varchar(256) DEFAULT NULL,
  `status` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_order_models_delete_at` (`delete_at`)
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_models`
--

LOCK TABLES `order_models` WRITE;
/*!40000 ALTER TABLE `order_models` DISABLE KEYS */;
INSERT INTO `order_models` VALUES (1,'2025-01-13 07:17:34',NULL,2,'cny',1,'test1@qq.com',0),(2,'2025-01-13 07:52:29',NULL,2,'cny',1,'test1@qq.com',0),(7,'2025-01-13 12:35:00',NULL,2,'cny',1,'test1@qq.com',2),(8,'2025-01-13 12:44:05',NULL,2,'cny',1,'test1@qq.com',2),(9,'2025-01-13 12:49:30',NULL,2,'cny',1,'test1@qq.com',2),(10,'2025-01-13 13:15:25',NULL,2,'cny',1,'test1@qq.com',2),(11,'2025-01-14 14:27:41',NULL,2,'cny',2,'test1@qq.com',2),(12,'2025-01-14 14:30:54',NULL,2,'cny',2,'test1@qq.com',2),(13,'2025-01-14 14:32:12',NULL,2,'cny',2,'test1@qq.com',2),(14,'2025-01-14 14:34:55',NULL,2,'cny',2,'test1@qq.com',2),(15,'2025-01-14 14:35:58',NULL,2,'cny',2,'test1@qq.com',2),(16,'2025-01-14 14:42:36',NULL,2,'cny',2,'test1@qq.com',2),(17,'2025-01-14 14:43:50',NULL,2,'cny',2,'test1@qq.com',1),(18,'2025-01-14 14:49:56',NULL,2,'cny',2,'test1@qq.com',1),(19,'2025-01-14 14:51:42',NULL,2,'cny',2,'test1@qq.com',2),(20,'2025-01-14 14:52:56',NULL,2,'cny',2,'test1@qq.com',2),(21,'2025-01-14 14:54:50',NULL,2,'cny',2,'test1@qq.com',1),(23,'2025-01-15 06:49:41',NULL,2,'cny',2,'test1@qq.com',1),(24,'2025-01-15 06:59:30',NULL,2,'cny',2,'test1@qq.com',1),(25,'2025-01-15 07:08:21',NULL,2,'cny',2,'test1@qq.com',1),(26,'2025-01-15 07:18:39',NULL,2,'cny',2,'test1@qq.com',2),(27,'2025-01-15 07:20:26',NULL,2,'cny',2,'test1@qq.com',2),(28,'2025-01-15 07:23:38',NULL,2,'cny',2,'test1@qq.com',1),(29,'2025-01-15 07:24:14',NULL,2,'cny',2,'test1@qq.com',1),(30,'2025-01-15 07:24:30',NULL,2,'cny',2,'test1@qq.com',1),(31,'2025-01-15 07:35:07',NULL,2,'cny',2,'test1@qq.com',1),(32,'2025-01-15 07:36:09',NULL,2,'cny',2,'test1@qq.com',1),(33,'2025-01-15 07:36:13',NULL,2,'cny',2,'test1@qq.com',2),(34,'2025-01-15 07:39:21',NULL,2,'cny',2,'test1@qq.com',1),(35,'2025-01-15 07:39:49',NULL,2,'cny',2,'test1@qq.com',1),(36,'2025-01-15 07:39:54',NULL,2,'cny',2,'test1@qq.com',2),(37,'2025-01-15 07:41:19',NULL,2,'cny',2,'test1@qq.com',2),(38,'2025-01-15 07:43:12',NULL,2,'cny',2,'test1@qq.com',1),(39,'2025-01-15 07:43:34',NULL,2,'cny',2,'test1@qq.com',1),(40,'2025-01-15 07:43:42',NULL,2,'cny',2,'test1@qq.com',1),(41,'2025-01-15 07:43:47',NULL,2,'cny',2,'test1@qq.com',2),(42,'2025-01-15 07:59:43',NULL,2,'cny',2,'test1@qq.com',1),(43,'2025-01-15 07:59:57',NULL,2,'cny',2,'test1@qq.com',2),(44,'2025-01-15 08:03:13',NULL,2,'cny',2,'test1@qq.com',2),(45,'2025-01-15 08:04:34',NULL,2,'cny',2,'test1@qq.com',2),(46,'2025-01-15 08:04:48',NULL,2,'cny',2,'test1@qq.com',2),(47,'2025-01-15 08:09:47',NULL,2,'cny',2,'test1@qq.com',1),(48,'2025-01-15 08:09:56',NULL,2,'cny',2,'test1@qq.com',1),(49,'2025-01-15 08:10:01',NULL,2,'cny',2,'test1@qq.com',1),(50,'2025-01-15 08:10:33',NULL,2,'cny',2,'test1@qq.com',1),(51,'2025-01-15 08:10:42',NULL,2,'cny',2,'test1@qq.com',2),(52,'2025-01-15 08:13:44',NULL,2,'cny',2,'test1@qq.com',2),(53,'2025-01-15 08:14:35',NULL,2,'cny',2,'test1@qq.com',2),(54,'2025-01-15 08:15:27',NULL,2,'cny',2,'test1@qq.com',2),(55,'2025-01-15 08:15:53',NULL,2,'cny',2,'test1@qq.com',2),(56,'2025-01-15 08:16:18',NULL,2,'cny',2,'test1@qq.com',2),(57,'2025-01-15 08:16:27',NULL,2,'cny',2,'test1@qq.com',2),(58,'2025-01-15 08:16:32',NULL,2,'cny',2,'test1@qq.com',2),(59,'2025-01-15 08:17:21',NULL,2,'cny',2,'test1@qq.com',2),(60,'2025-01-15 08:19:20',NULL,2,'cny',2,'test1@qq.com',2),(61,'2025-01-15 08:19:29',NULL,2,'cny',2,'test1@qq.com',2),(62,'2025-01-15 08:21:22',NULL,2,'cny',2,'test1@qq.com',2),(63,'2025-01-15 08:21:30',NULL,2,'cny',2,'test1@qq.com',1),(64,'2025-01-15 09:04:15',NULL,2,'cny',2,'test1@qq.com',2),(65,'2025-01-15 09:08:25',NULL,2,'cny',2,'test1@qq.com',1),(66,'2025-01-15 09:08:45',NULL,2,'cny',2,'test1@qq.com',1),(67,'2025-01-15 09:08:50',NULL,2,'cny',2,'test1@qq.com',2),(68,'2025-01-15 09:09:08',NULL,2,'cny',2,'test1@qq.com',1),(69,'2025-01-15 09:09:14',NULL,2,'cny',2,'test1@qq.com',1),(70,'2025-01-15 09:09:18',NULL,2,'cny',2,'test1@qq.com',1),(71,'2025-01-15 09:09:22',NULL,2,'cny',2,'test1@qq.com',1),(72,'2025-01-15 09:09:30',NULL,2,'cny',2,'test1@qq.com',2),(73,'2025-01-15 11:50:02',NULL,2,'cny',2,'test1@qq.com',1);
/*!40000 ALTER TABLE `order_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_log_models`
--

DROP TABLE IF EXISTS `payment_log_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payment_log_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_at` datetime(3) DEFAULT NULL,
  `transaction_id` bigint unsigned NOT NULL,
  `action` tinyint NOT NULL,
  `message` varchar(256) DEFAULT NULL,
  `status` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_payment_log_models_delete_at` (`delete_at`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_log_models`
--

LOCK TABLES `payment_log_models` WRITE;
/*!40000 ALTER TABLE `payment_log_models` DISABLE KEYS */;
INSERT INTO `payment_log_models` VALUES (1,'2025-01-14 14:35:58',NULL,389734749953,0,'Payment Created',0),(2,'2025-01-14 14:42:36',NULL,1058206144257,0,'Payment Created',0),(3,'2025-01-14 14:43:51',NULL,1183246735105,0,'Payment Created',0),(4,'2025-01-14 14:49:56',NULL,1795850000129,0,'Payment Created',0),(5,'2025-01-14 14:51:42',NULL,1974728677121,0,'Payment Created',0),(6,'2025-01-14 14:52:56',NULL,13908334337,0,'Payment Created',0),(7,'2025-01-14 14:54:50',NULL,9630144257,0,'Payment Created',0),(8,'2025-01-15 06:49:41',NULL,567992669953,0,'Payment Created',0),(9,'2025-01-15 06:59:30',NULL,47043335937,0,'Payment failed due to insufficient funds or network error',2),(10,'2025-01-15 07:08:21',NULL,937594738433,0,'Payment failed due to insufficient funds or network error',2),(11,'2025-01-15 07:18:39',NULL,30316451585,0,'Payment failed due to insufficient funds or network error',2),(12,'2025-01-15 07:20:26',NULL,210889627393,0,'Payment failed due to insufficient funds or network error',2),(13,'2025-01-15 07:23:38',NULL,24947742465,0,'Payment processed successfully',1),(14,'2025-01-15 07:24:14',NULL,84976621313,0,'Payment processed successfully',1),(15,'2025-01-15 07:24:30',NULL,110897420033,0,'Payment processed successfully',1),(16,'2025-01-15 07:35:07',NULL,28504512257,0,'Payment processed successfully',1),(17,'2025-01-15 07:36:09',NULL,131382400769,0,'Payment processed successfully',1),(18,'2025-01-15 07:36:13',NULL,137707411201,0,'Payment failed due to insufficient funds or network error',2),(19,'2025-01-15 07:39:21',NULL,8774506241,0,'Payment processed successfully',1),(20,'2025-01-15 07:39:49',NULL,56270804737,0,'Payment processed successfully',1),(21,'2025-01-15 07:39:54',NULL,64458086145,0,'Payment failed due to insufficient funds or network error',2),(22,'2025-01-15 07:41:19',NULL,207131531009,0,'Payment failed due to insufficient funds or network error',2),(23,'2025-01-15 07:43:12',NULL,62646146817,0,'Payment processed successfully',1),(24,'2025-01-15 07:43:34',NULL,100109670145,0,'Payment processed successfully',1),(25,'2025-01-15 07:43:42',NULL,112843577089,0,'Payment processed successfully',1),(26,'2025-01-15 07:43:47',NULL,121550952193,0,'Payment failed due to insufficient funds or network error',2),(27,'2025-01-15 07:59:43',NULL,30098347777,0,'Payment processed successfully',1),(28,'2025-01-15 07:59:57',NULL,54626637569,0,'Payment failed due to insufficient funds or network error',2),(29,'2025-01-15 08:03:13',NULL,382671542017,0,'Payment failed due to insufficient funds or network error',2),(30,'2025-01-15 08:04:34',NULL,518114006785,0,'Payment processed successfully',1),(31,'2025-01-15 08:04:48',NULL,541652440833,0,'Payment failed due to insufficient funds or network error',2),(32,'2025-01-15 08:09:47',NULL,1044029396737,0,'Payment processed successfully',1),(33,'2025-01-15 08:09:56',NULL,1059850311425,0,'Payment processed successfully',1),(34,'2025-01-15 08:10:01',NULL,1068004038401,0,'Payment processed successfully',1),(35,'2025-01-15 08:10:33',NULL,1121372362497,0,'Payment processed successfully',1),(36,'2025-01-15 08:10:42',NULL,1136840955649,0,'Payment failed due to insufficient funds or network error',2),(37,'2025-01-15 08:13:44',NULL,1441783633665,0,'Payment processed successfully',1),(38,'2025-01-15 08:14:35',NULL,15871268609,0,'Payment failed due to insufficient funds or network error',2),(39,'2025-01-15 08:15:27',NULL,102240376577,0,'Payment failed due to insufficient funds or network error',2),(40,'2025-01-15 08:15:53',NULL,146515449601,0,'Payment failed due to insufficient funds or network error',2),(41,'2025-01-15 08:16:18',NULL,188928251649,0,'Payment failed due to insufficient funds or network error',2),(42,'2025-01-15 08:16:27',NULL,203910305537,0,'Payment failed due to insufficient funds or network error',2),(43,'2025-01-15 08:16:32',NULL,212231804673,0,'Payment failed due to insufficient funds or network error',2),(44,'2025-01-15 08:17:21',NULL,293987178241,0,'Payment failed due to insufficient funds or network error',2),(45,'2025-01-15 08:19:20',NULL,493032068865,0,'Payment failed due to insufficient funds or network error',2),(46,'2025-01-15 08:19:29',NULL,509742176001,0,'Payment failed due to insufficient funds or network error',2),(47,'2025-01-15 08:21:22',NULL,76923557633,0,'Payment failed due to insufficient funds or network error',2),(48,'2025-01-15 08:21:30',NULL,89858791169,0,'Payment processed successfully',1),(49,'2025-01-15 09:00:27',NULL,0,2,'Payment Canceled',3),(50,'2025-01-15 09:02:18',NULL,0,2,'Payment Canceled',3),(51,'2025-01-15 09:04:15',NULL,245551355649,0,'Payment processed successfully',1),(52,'2025-01-15 09:08:25',NULL,666290378497,0,'Payment processed successfully',1),(53,'2025-01-15 09:08:45',NULL,699526043393,0,'Payment processed successfully',1),(54,'2025-01-15 09:08:50',NULL,707696547585,0,'Payment failed due to insufficient funds or network error',2),(55,'2025-01-15 09:09:08',NULL,737895536385,0,'Payment processed successfully',1),(56,'2025-01-15 09:09:14',NULL,747357886209,0,'Payment processed successfully',1),(57,'2025-01-15 09:09:18',NULL,755142514433,0,'Payment processed successfully',1),(58,'2025-01-15 09:09:22',NULL,761987618561,0,'Payment processed successfully',1),(59,'2025-01-15 09:09:30',NULL,774587307777,0,'Payment failed due to insufficient funds or network error',2);
/*!40000 ALTER TABLE `payment_log_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_models`
--

DROP TABLE IF EXISTS `payment_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payment_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_at` datetime(3) DEFAULT NULL,
  `user_id` int unsigned NOT NULL,
  `order_id` int unsigned NOT NULL,
  `transaction_id` bigint unsigned NOT NULL,
  `amount` double NOT NULL,
  `currency` varchar(16) DEFAULT NULL,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_payment_models_delete_at` (`delete_at`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_models`
--

LOCK TABLES `payment_models` WRITE;
/*!40000 ALTER TABLE `payment_models` DISABLE KEYS */;
INSERT INTO `payment_models` VALUES (1,'2025-01-14 14:35:58',NULL,2,15,0,36369,'','2025-01-14 14:35:58',3),(2,'2025-01-14 14:42:36',NULL,2,16,0,36369,'','2025-01-14 14:42:36',3),(3,'2025-01-14 14:43:51',NULL,2,17,0,36369,'','2025-01-14 14:43:51',3),(4,'2025-01-14 14:49:56',NULL,2,18,0,36369,'','2025-01-14 14:49:56',3),(5,'2025-01-14 14:51:42',NULL,2,19,0,36369,'','2025-01-14 14:51:42',3),(6,'2025-01-14 14:52:56',NULL,2,20,0,36369,'','2025-01-14 14:52:56',3),(7,'2025-01-14 14:54:50',NULL,2,21,0,36369,'','2025-01-14 14:54:50',0),(8,'2025-01-15 06:49:41',NULL,2,23,0,36369,'','2025-01-15 06:49:41',0),(9,'2025-01-15 06:59:30',NULL,2,24,0,36369,'','2025-01-15 06:59:30',2),(10,'2025-01-15 07:08:21',NULL,2,25,0,36369,'','2025-01-15 07:08:21',2),(11,'2025-01-15 07:18:39',NULL,2,26,0,36369,'','2025-01-15 07:18:39',2),(12,'2025-01-15 07:20:26',NULL,2,27,0,36369,'','2025-01-15 07:20:26',2),(13,'2025-01-15 07:23:38',NULL,2,28,0,36369,'','2025-01-15 07:23:38',1),(14,'2025-01-15 07:24:14',NULL,2,29,0,36369,'','2025-01-15 07:24:14',1),(15,'2025-01-15 07:24:30',NULL,2,30,0,36369,'','2025-01-15 07:24:30',1),(16,'2025-01-15 07:35:07',NULL,2,31,28504512257,36369,'','2025-01-15 07:35:07',1),(17,'2025-01-15 07:36:09',NULL,2,32,131382400769,36369,'','2025-01-15 07:36:09',1),(18,'2025-01-15 07:36:13',NULL,2,33,137707411201,36369,'','2025-01-15 07:36:13',2),(19,'2025-01-15 07:39:21',NULL,2,34,8774506241,36369,'','2025-01-15 07:39:21',1),(20,'2025-01-15 07:39:49',NULL,2,35,56270804737,36369,'','2025-01-15 07:39:49',1),(21,'2025-01-15 07:39:54',NULL,2,36,64458086145,36369,'','2025-01-15 07:39:54',2),(22,'2025-01-15 07:41:19',NULL,2,37,207131531009,36369,'','2025-01-15 07:41:19',2),(23,'2025-01-15 07:43:12',NULL,2,38,62646146817,36369,'','2025-01-15 07:43:12',1),(24,'2025-01-15 07:43:34',NULL,2,39,100109670145,36369,'','2025-01-15 07:43:34',1),(25,'2025-01-15 07:43:42',NULL,2,40,112843577089,36369,'','2025-01-15 07:43:42',1),(26,'2025-01-15 07:43:47',NULL,2,41,121550952193,36369,'','2025-01-15 07:43:47',2),(27,'2025-01-15 07:59:43',NULL,2,42,30098347777,36369,'','2025-01-15 07:59:43',1),(28,'2025-01-15 07:59:57',NULL,2,43,54626637569,36369,'','2025-01-15 07:59:57',2),(29,'2025-01-15 08:03:13',NULL,2,44,382671542017,36369,'','2025-01-15 08:03:13',2),(30,'2025-01-15 08:04:34',NULL,2,45,518114006785,36369,'','2025-01-15 08:04:34',1),(31,'2025-01-15 08:04:48',NULL,2,46,541652440833,36369,'','2025-01-15 08:04:48',2),(32,'2025-01-15 08:09:47',NULL,2,47,1044029396737,36369,'','2025-01-15 08:09:47',1),(33,'2025-01-15 08:09:56',NULL,2,48,1059850311425,36369,'','2025-01-15 08:09:56',1),(34,'2025-01-15 08:10:01',NULL,2,49,1068004038401,36369,'','2025-01-15 08:10:01',1),(35,'2025-01-15 08:10:33',NULL,2,50,1121372362497,36369,'','2025-01-15 08:10:33',1),(36,'2025-01-15 08:10:42',NULL,2,51,1136840955649,36369,'','2025-01-15 08:10:42',2),(37,'2025-01-15 08:13:44',NULL,2,52,1441783633665,36369,'','2025-01-15 08:13:44',1),(38,'2025-01-15 08:14:35',NULL,2,53,15871268609,36369,'','2025-01-15 08:14:35',2),(39,'2025-01-15 08:15:27',NULL,2,54,102240376577,36369,'','2025-01-15 08:15:27',2),(40,'2025-01-15 08:15:53',NULL,2,55,146515449601,36369,'','2025-01-15 08:15:53',2),(41,'2025-01-15 08:16:18',NULL,2,56,188928251649,36369,'','2025-01-15 08:16:18',2),(42,'2025-01-15 08:16:27',NULL,2,57,203910305537,36369,'','2025-01-15 08:16:27',2),(43,'2025-01-15 08:16:32',NULL,2,58,212231804673,36369,'','2025-01-15 08:16:32',2),(44,'2025-01-15 08:17:21',NULL,2,59,293987178241,36369,'','2025-01-15 08:17:21',2),(45,'2025-01-15 08:19:20',NULL,2,60,493032068865,36369,'','2025-01-15 08:19:20',2),(46,'2025-01-15 08:19:29',NULL,2,61,509742176001,36369,'','2025-01-15 08:19:29',2),(47,'2025-01-15 08:21:22',NULL,2,62,76923557633,36369,'','2025-01-15 08:21:22',2),(48,'2025-01-15 08:21:30',NULL,2,63,89858791169,36369,'','2025-01-15 08:21:30',1),(49,'2025-01-15 09:04:15',NULL,2,64,245551355649,36369,'','2025-01-15 09:04:15',1),(50,'2025-01-15 09:08:25',NULL,2,65,666290378497,36369,'','2025-01-15 09:08:25',1),(51,'2025-01-15 09:08:45',NULL,2,66,699526043393,36369,'','2025-01-15 09:08:45',1),(52,'2025-01-15 09:08:50',NULL,2,67,707696547585,36369,'','2025-01-15 09:08:50',2),(53,'2025-01-15 09:09:08',NULL,2,68,737895536385,36369,'','2025-01-15 09:09:08',1),(54,'2025-01-15 09:09:14',NULL,2,69,747357886209,36369,'','2025-01-15 09:09:14',1),(55,'2025-01-15 09:09:18',NULL,2,70,755142514433,36369,'','2025-01-15 09:09:18',1),(56,'2025-01-15 09:09:22',NULL,2,71,761987618561,36369,'','2025-01-15 09:09:22',1),(57,'2025-01-15 09:09:30',NULL,2,72,774587307777,36369,'','2025-01-15 09:09:30',2);
/*!40000 ALTER TABLE `payment_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_models`
--

DROP TABLE IF EXISTS `product_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_at` datetime(3) DEFAULT NULL,
  `name` varchar(64) NOT NULL,
  `description` varchar(256) NOT NULL,
  `picture` varchar(256) DEFAULT NULL,
  `price` double NOT NULL DEFAULT '0',
  `stock` bigint NOT NULL DEFAULT '0',
  `status` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_product_models_delete_at` (`delete_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_models`
--

LOCK TABLES `product_models` WRITE;
/*!40000 ALTER TABLE `product_models` DISABLE KEYS */;
INSERT INTO `product_models` VALUES (1,'2025-01-09 13:45:45',NULL,'苹果14','这个一个又大又圆的苹果14，5块钱一个！','www.test1.com/pic1.jpg',12123,70,0),(2,'2025-01-09 13:46:52',NULL,'苹果15','这个一个又大又圆的苹果15，55块钱一个！','www.test1.com/pic1.jpg',12123,100,0),(3,'2025-01-09 13:46:59',NULL,'苹果16','这个一个又大又圆的苹果16，55块钱一个！','www.test1.com/pic1.jpg',12123,1,0),(4,'2025-01-09 13:47:16',NULL,'苹果13','这个一个又大又圆的苹果13，55块钱一个！！','www.test1.com/pic1.jpg',12123,10,0),(5,'2025-01-09 13:47:30','2025-01-09 22:02:21.472','苹果16','这个一个又大又圆的苹果12，55块钱一个！','www.test1.com/pic1.jpg',12123,100,0),(6,'2025-01-15 16:13:30',NULL,'苹果16','这个一个又大又圆的苹果16，55块钱一个！','www.test1.com/pic1.jpg',12123,100,0),(7,'2025-01-15 16:15:01',NULL,'苹果17','这个一个又大又圆的苹果17，55块钱一个！','www.test1.com/pic1.jpg',55,200,0);
/*!40000 ALTER TABLE `product_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_models`
--

DROP TABLE IF EXISTS `role_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_role_models_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_models`
--

LOCK TABLES `role_models` WRITE;
/*!40000 ALTER TABLE `role_models` DISABLE KEYS */;
INSERT INTO `role_models` VALUES (1,'admin','2025-02-11 05:37:25'),(2,'merchant','2025-02-11 05:37:25'),(3,'user','2025-02-11 05:37:25');
/*!40000 ALTER TABLE `role_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_models`
--

DROP TABLE IF EXISTS `user_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_at` datetime(3) DEFAULT NULL,
  `password` varchar(64) NOT NULL,
  `status` tinyint NOT NULL DEFAULT '0',
  `email` varchar(64) DEFAULT NULL,
  `verified` tinyint(1) DEFAULT '0',
  `currency` varchar(16) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_models_delete_at` (`delete_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_models`
--

LOCK TABLES `user_models` WRITE;
/*!40000 ALTER TABLE `user_models` DISABLE KEYS */;
INSERT INTO `user_models` VALUES (2,'2024-12-24 13:11:24',NULL,'ef8364020560a703cfc7aebebcd0b62d1bfea7d4a841eb8964cfbcda2ba85dd5',0,'test1@qq.com',0,'cny'),(3,'2025-02-11 06:45:16',NULL,'ef8364020560a703cfc7aebebcd0b62d1bfea7d4a841eb8964cfbcda2ba85dd5',0,'test2@qq.com',0,''),(4,'2025-02-12 02:28:37',NULL,'ef8364020560a703cfc7aebebcd0b62d1bfea7d4a841eb8964cfbcda2ba85dd5',0,'test3@qq.com',0,''),(5,'2025-02-12 02:29:35',NULL,'ef8364020560a703cfc7aebebcd0b62d1bfea7d4a841eb8964cfbcda2ba85dd5',0,'test4@qq.com',0,''),(6,'2025-02-12 02:32:11',NULL,'ef8364020560a703cfc7aebebcd0b62d1bfea7d4a841eb8964cfbcda2ba85dd5',0,'test5@qq.com',0,'');
/*!40000 ALTER TABLE `user_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_roles`
--

DROP TABLE IF EXISTS `user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_roles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `role_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_roles_user_id` (`user_id`),
  KEY `idx_user_roles_role_id` (`role_id`),
  CONSTRAINT `fk_user_roles_role` FOREIGN KEY (`role_id`) REFERENCES `role_models` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_user_roles_user` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT INTO `user_roles` VALUES (1,2,3);
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `white_list_models`
--

DROP TABLE IF EXISTS `white_list_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `white_list_models` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `path` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `white_list_models`
--

LOCK TABLES `white_list_models` WRITE;
/*!40000 ALTER TABLE `white_list_models` DISABLE KEYS */;
INSERT INTO `white_list_models` VALUES (1,'/merchant'),(2,'/api/product/detail'),(3,'/api/user/login'),(4,'/api/user/register'),(5,'/api/user/logout'),(6,'/api/user/verify-access-token'),(7,'/api/user/refresh-access-token');
/*!40000 ALTER TABLE `white_list_models` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-12  8:07:12

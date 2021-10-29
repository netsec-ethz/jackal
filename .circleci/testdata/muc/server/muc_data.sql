-- MariaDB dump 10.19  Distrib 10.6.4-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: jackal
-- ------------------------------------------------------
-- Server version	10.6.4-MariaDB

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
-- Table structure for table `blocklist_items`
--

DROP TABLE IF EXISTS `blocklist_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blocklist_items` (
  `username` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `jid` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`username`,`jid`),
  KEY `i_blocklist_items_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blocklist_items`
--

LOCK TABLES `blocklist_items` WRITE;
/*!40000 ALTER TABLE `blocklist_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `blocklist_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `capabilities`
--

DROP TABLE IF EXISTS `capabilities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `capabilities` (
  `node` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ver` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `features` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`node`,`ver`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `capabilities`
--

LOCK TABLES `capabilities` WRITE;
/*!40000 ALTER TABLE `capabilities` DISABLE KEYS */;
/*!40000 ALTER TABLE `capabilities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `occupants`
--

DROP TABLE IF EXISTS `occupants`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `occupants` (
  `occupant_jid` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `bare_jid` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `affiliation` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `role` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`occupant_jid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `occupants`
--

LOCK TABLES `occupants` WRITE;
/*!40000 ALTER TABLE `occupants` DISABLE KEYS */;
/*!40000 ALTER TABLE `occupants` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `offline_messages`
--

DROP TABLE IF EXISTS `offline_messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `offline_messages` (
  `username` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL,
  KEY `i_offline_messages_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `offline_messages`
--

LOCK TABLES `offline_messages` WRITE;
/*!40000 ALTER TABLE `offline_messages` DISABLE KEYS */;
/*!40000 ALTER TABLE `offline_messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `presences`
--

DROP TABLE IF EXISTS `presences`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `presences` (
  `username` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `domain` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `resource` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `presence` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `node` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ver` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `allocation_id` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`username`,`domain`,`resource`),
  KEY `i_presences_username_domain` (`username`,`domain`),
  KEY `i_presences_domain_resource` (`domain`,`resource`),
  KEY `i_presences_allocation_id` (`allocation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `presences`
--

LOCK TABLES `presences` WRITE;
/*!40000 ALTER TABLE `presences` DISABLE KEYS */;
/*!40000 ALTER TABLE `presences` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `private_storage`
--

DROP TABLE IF EXISTS `private_storage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `private_storage` (
  `username` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `namespace` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`username`,`namespace`),
  KEY `i_private_storage_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `private_storage`
--

LOCK TABLES `private_storage` WRITE;
/*!40000 ALTER TABLE `private_storage` DISABLE KEYS */;
/*!40000 ALTER TABLE `private_storage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pubsub_affiliations`
--

DROP TABLE IF EXISTS `pubsub_affiliations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pubsub_affiliations` (
  `node_id` bigint(20) NOT NULL,
  `jid` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `affiliation` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  UNIQUE KEY `i_pubsub_affiliations_node_id_jid` (`node_id`,`jid`(512)),
  KEY `i_pubsub_affiliations_jid` (`jid`(512))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pubsub_affiliations`
--

LOCK TABLES `pubsub_affiliations` WRITE;
/*!40000 ALTER TABLE `pubsub_affiliations` DISABLE KEYS */;
/*!40000 ALTER TABLE `pubsub_affiliations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pubsub_items`
--

DROP TABLE IF EXISTS `pubsub_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pubsub_items` (
  `node_id` bigint(20) NOT NULL,
  `item_id` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `publisher` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  UNIQUE KEY `i_pubsub_items_node_id_item_id` (`node_id`,`item_id`(36)),
  KEY `i_pubsub_items_item_id` (`item_id`(36)),
  KEY `i_pubsub_items_node_id_created_at` (`node_id`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pubsub_items`
--

LOCK TABLES `pubsub_items` WRITE;
/*!40000 ALTER TABLE `pubsub_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `pubsub_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pubsub_node_options`
--

DROP TABLE IF EXISTS `pubsub_node_options`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pubsub_node_options` (
  `node_id` bigint(20) NOT NULL,
  `name` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  KEY `i_pubsub_node_options_node_id` (`node_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pubsub_node_options`
--

LOCK TABLES `pubsub_node_options` WRITE;
/*!40000 ALTER TABLE `pubsub_node_options` DISABLE KEYS */;
/*!40000 ALTER TABLE `pubsub_node_options` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pubsub_nodes`
--

DROP TABLE IF EXISTS `pubsub_nodes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pubsub_nodes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `host` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `i_pubsub_nodes_host_name` (`host`(256),`name`(512)),
  KEY `i_pubsub_nodes_host` (`host`(256))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pubsub_nodes`
--

LOCK TABLES `pubsub_nodes` WRITE;
/*!40000 ALTER TABLE `pubsub_nodes` DISABLE KEYS */;
/*!40000 ALTER TABLE `pubsub_nodes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pubsub_subscriptions`
--

DROP TABLE IF EXISTS `pubsub_subscriptions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pubsub_subscriptions` (
  `node_id` bigint(20) NOT NULL,
  `subid` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `jid` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `subscription` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  UNIQUE KEY `i_pubsub_subscriptions_node_id_jid` (`node_id`,`jid`(512)),
  KEY `i_pubsub_subscriptions_jid` (`jid`(512))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pubsub_subscriptions`
--

LOCK TABLES `pubsub_subscriptions` WRITE;
/*!40000 ALTER TABLE `pubsub_subscriptions` DISABLE KEYS */;
/*!40000 ALTER TABLE `pubsub_subscriptions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `resources`
--

DROP TABLE IF EXISTS `resources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `resources` (
  `occupant_jid` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `resource` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`occupant_jid`,`resource`),
  KEY `i_occupant_jid` (`occupant_jid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `resources`
--

LOCK TABLES `resources` WRITE;
/*!40000 ALTER TABLE `resources` DISABLE KEYS */;
/*!40000 ALTER TABLE `resources` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rooms`
--

DROP TABLE IF EXISTS `rooms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rooms` (
  `room_jid` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `subject` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `language` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `locked` tinyint(1) NOT NULL,
  `occupants_online` int(11) NOT NULL,
  PRIMARY KEY (`room_jid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rooms`
--

LOCK TABLES `rooms` WRITE;
/*!40000 ALTER TABLE `rooms` DISABLE KEYS */;
/*!40000 ALTER TABLE `rooms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rooms_config`
--

DROP TABLE IF EXISTS `rooms_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rooms_config` (
  `room_jid` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `public` tinyint(1) NOT NULL,
  `persistent` tinyint(1) NOT NULL,
  `pwd_protected` tinyint(1) NOT NULL,
  `password` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `open` tinyint(1) NOT NULL,
  `moderated` tinyint(1) NOT NULL,
  `allow_invites` tinyint(1) NOT NULL,
  `max_occupants` int(11) NOT NULL,
  `allow_subj_change` tinyint(1) NOT NULL,
  `non_anonymous` tinyint(1) NOT NULL,
  `can_send_pm` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `can_get_member_list` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`room_jid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rooms_config`
--

LOCK TABLES `rooms_config` WRITE;
/*!40000 ALTER TABLE `rooms_config` DISABLE KEYS */;
/*!40000 ALTER TABLE `rooms_config` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rooms_invites`
--

DROP TABLE IF EXISTS `rooms_invites`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rooms_invites` (
  `room_jid` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_jid` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`room_jid`,`user_jid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rooms_invites`
--

LOCK TABLES `rooms_invites` WRITE;
/*!40000 ALTER TABLE `rooms_invites` DISABLE KEYS */;
/*!40000 ALTER TABLE `rooms_invites` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rooms_users`
--

DROP TABLE IF EXISTS `rooms_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rooms_users` (
  `room_jid` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_jid` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `occupant_jid` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`room_jid`,`user_jid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rooms_users`
--

LOCK TABLES `rooms_users` WRITE;
/*!40000 ALTER TABLE `rooms_users` DISABLE KEYS */;
/*!40000 ALTER TABLE `rooms_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roster_groups`
--

DROP TABLE IF EXISTS `roster_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roster_groups` (
  `username` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `jid` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `group` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  KEY `i_roster_groups_username_jid` (`username`,`jid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roster_groups`
--

LOCK TABLES `roster_groups` WRITE;
/*!40000 ALTER TABLE `roster_groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `roster_groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roster_items`
--

DROP TABLE IF EXISTS `roster_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roster_items` (
  `username` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `jid` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `subscription` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `groups` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `ask` tinyint(1) NOT NULL,
  `ver` int(11) NOT NULL DEFAULT 0,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`username`,`jid`),
  KEY `i_roster_items_username` (`username`),
  KEY `i_roster_items_jid` (`jid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roster_items`
--

LOCK TABLES `roster_items` WRITE;
/*!40000 ALTER TABLE `roster_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `roster_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roster_notifications`
--

DROP TABLE IF EXISTS `roster_notifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roster_notifications` (
  `contact` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `jid` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `elements` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`contact`,`jid`),
  KEY `i_roster_notifications_jid` (`jid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roster_notifications`
--

LOCK TABLES `roster_notifications` WRITE;
/*!40000 ALTER TABLE `roster_notifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `roster_notifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roster_versions`
--

DROP TABLE IF EXISTS `roster_versions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roster_versions` (
  `username` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ver` int(11) NOT NULL DEFAULT 0,
  `last_deletion_ver` int(11) NOT NULL DEFAULT 0,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roster_versions`
--

LOCK TABLES `roster_versions` WRITE;
/*!40000 ALTER TABLE `roster_versions` DISABLE KEYS */;
/*!40000 ALTER TABLE `roster_versions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `username` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_presence` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_presence_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('alice','asdf','<presence from=\"alice@muc_server.xmpp/profanity\" to=\"alice@muc_server.xmpp\" type=\"unavailable\"/>','2021-10-24 18:42:58','2021-10-24 18:42:58','2021-10-24 18:42:58'),('bob','asdf','<presence from=\"bob@muc_server.xmpp/profanity\" to=\"bob@muc_server.xmpp\" type=\"unavailable\"/>','2021-10-24 18:42:58','2021-10-24 18:42:58','2021-10-24 18:42:58'),('carol','asdf','<presence from=\"carol@muc_server.xmpp/profanity\" to=\"carol@muc_server.xmpp\" type=\"unavailable\"/>','2021-10-24 18:42:58','2021-10-24 18:42:58','2021-10-24 18:42:58');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vcards`
--

DROP TABLE IF EXISTS `vcards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vcards` (
  `username` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `vcard` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vcards`
--

LOCK TABLES `vcards` WRITE;
/*!40000 ALTER TABLE `vcards` DISABLE KEYS */;
/*!40000 ALTER TABLE `vcards` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-10-25 16:01:21

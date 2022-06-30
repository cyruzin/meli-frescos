CREATE DATABASE IF NOT EXISTS `bootcamp`;

USE `bootcamp`;

DROP TABLE IF EXISTS `sections`;

CREATE TABLE `sections` (
  id                  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  section_number      INT UNSIGNED NOT NULL,
  current_temperature INT NOT NULL,
  minimum_temperature INT NOT NULL,
  current_capacity    INT UNSIGNED NOT NULL,
  minimum_capacity    INT UNSIGNED NOT NULL,
  maximum_capacity    INT UNSIGNED NOT NULL,
  warehouse_id        INT UNSIGNED NOT NULL,
  product_type_id     INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `sections` WRITE;

UNLOCK TABLES;
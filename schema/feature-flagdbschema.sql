CREATE DATABASE IF NOT EXISTS `feature-flag-db`;

CREATE TABLE IF NOT EXISTS `feature`(
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `created_at` DATE NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (name));

CREATE TABLE IF NOT EXISTS `environment` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `feature_id` INT NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `enabled` BOOLEAN NOT NULL,
  `created_at` DATE NOT NULL,
  PRIMARY KEY (id));

CREATE TABLE IF NOT EXISTS `environment_history` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `feature_id` INT NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `enabled` BOOLEAN NOT NULL,
  `created_at` DATE NOT NULL,
  `timestamp` DATE NOT NULL,
  PRIMARY KEY (id));

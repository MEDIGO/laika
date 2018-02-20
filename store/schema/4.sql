-- +migrate Up

CREATE TABLE IF NOT EXISTS `events` (
  `id`            INT NOT NULL AUTO_INCREMENT,
  `time`          TIMESTAMP NOT NULL,
  `type`          VARCHAR(255) NOT NULL,
  `data`          MEDIUMTEXT NOT NULL,

  PRIMARY KEY (`id`),
  KEY(`time`),
  KEY(`type`)
) DEFAULT CHARSET=utf8;

-- +migrate Down

DROP TABLE IF EXISTS `events`;

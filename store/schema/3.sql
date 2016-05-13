-- +migrate Up

CREATE TABLE IF NOT EXISTS `user`(
  `id`            INT NOT NULL AUTO_INCREMENT,
  `username`      VARCHAR(255) NOT NULL,
  `password_hash` VARCHAR(255) NOT NULL,
  `created_at`    DATE NOT NULL,
  `updated_at`    DATE,

  PRIMARY KEY (id),
  UNIQUE (username)
);

-- +migrate Down

DROP TABLE IF EXISTS user;

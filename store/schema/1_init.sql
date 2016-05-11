-- +migrate Up

CREATE TABLE IF NOT EXISTS `feature`(
  `id`         INT NOT NULL AUTO_INCREMENT,
  `name`       VARCHAR(255) NOT NULL,
  `created_at` DATE NOT NULL,

  PRIMARY KEY (id),
  UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS `environment` (
  `id`         INT NOT NULL AUTO_INCREMENT,
  `name`       VARCHAR(255) NOT NULL,
  `created_at` DATE NOT NULL,

  PRIMARY KEY (id),
  UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS `feature_status` (
  `id`             INT NOT NULL AUTO_INCREMENT,
  `feature_id`     INT NOT NULL,
  `environment_id` INT NOT NULL,
  `enabled`        BOOLEAN NOT NULL,
  `created_at`     DATE NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (feature_id) REFERENCES feature(id),
  FOREIGN KEY (environment_id) REFERENCES environment(id)
);

CREATE TABLE IF NOT EXISTS `feature_status_history` (
  `id`                INT NOT NULL AUTO_INCREMENT,
  `feature_id`        INT NOT NULL,
  `environment_id`    INT NOT NULL,
  `feature_status_id` INT NOT NULL,
  `enabled`           BOOLEAN NOT NULL,
  `created_at`        DATE NOT NULL,
  `timestamp`         DATE NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (feature_id) REFERENCES feature(id),
  FOREIGN KEY (environment_id) REFERENCES environment(id),
  FOREIGN KEY (feature_status_id) REFERENCES feature_status(id)
);

-- +migrate Down

DROP TABLE IF EXISTS feature_status_history;
DROP TABLE IF EXISTS feature_status;
DROP TABLE IF EXISTS environment;
DROP TABLE IF EXISTS feature;

-- +migrate Up

ALTER TABLE `feature_status_history` DROP COLUMN `timestamp`;

-- +migrate Down

ALTER TABLE `feature_status_history` ADD COLUMN `timestamp` DATE NOT NULL;

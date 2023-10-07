-- migrate:up
ALTER TABLE `documentations` ADD `name` varchar(36) NOT NULL AFTER `branch_id`;
ALTER TABLE `documentations` ADD `description` varchar(36) NOT NULL AFTER `location`;

-- migrate:down
ALTER TABLE `documentations` DROP COLUMN `name`;
ALTER TABLE `documentations` DROP COLUMN `description`;



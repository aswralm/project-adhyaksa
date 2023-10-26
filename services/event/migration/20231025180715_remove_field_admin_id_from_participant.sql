-- migrate:up
ALTER TABLE `participants` DROP COLUMN `admin_id`;

-- migrate:down
ALTER TABLE `participants` ADD `admin_id` varchar(36) NOT NULL AFTER `user_id`;

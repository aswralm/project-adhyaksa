-- migrate:up
ALTER TABLE `photos` ADD `public_id` varchar(36) NOT NULL;

-- migrate:down
ALTER TABLE `photos`DROP COLUMN `public_id`;

-- migrate:up
ALTER TABLE `events`DROP COLUMN `desctiption`;
ALTER TABLE `events` ADD `description` varchar(300) NOT NULL;

-- migrate:down
ALTER TABLE `events`DROP COLUMN `description`;

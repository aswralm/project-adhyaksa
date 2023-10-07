-- migrate:up
ALTER TABLE `photos` MODIFY COLUMN `url` VARCHAR(100);

-- migrate:down


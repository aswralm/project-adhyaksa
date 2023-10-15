-- migrate:up
ALTER TABLE `branchs` RENAME TO `branches`;

-- migrate:down
ALTER TABLE `branches` RENAME TO `branchs`;

-- migrate:up
CREATE TABLE `branchs` (
  `id` varchar(36) PRIMARY KEY NOT NULL,
  `name` varchar(36) NOT NULL,
  `address` varchar(36) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp NULL,
  `deleted_at` timestamp NULL
);

CREATE TABLE `events` (
  `id` varchar(36) PRIMARY KEY NOT NULL,
  `admin_id` varchar(36),
  `branch_id` varchar(36) NOT NULL,
  `name` varchar(36) NOT NULL,
  `start_time` timestamp NULL,
  `end_time` timestamp NULL,
  `location` varchar(36) NOT NULL,
  `desctiption` varchar(100) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp NULL,
  `deleted_at` timestamp NULL
);

CREATE TABLE `documentations` (
  `id` varchar(36) PRIMARY KEY NOT NULL,
  `admin_id` varchar(36),
  `branch_id` varchar(36) NOT NULL,
  `date` timestamp NOT NULL,
  `participant` int NOT NULL,
  `location` varchar(36) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp NULL,
  `deleted_at` timestamp NULL
);

CREATE TABLE `participants` (
  `id` varchar(36) PRIMARY KEY NOT NULL,
  `user_id` varchar(36),
  `admin_id` varchar(36),
  `event_id` varchar(36) NOT NULL,
  `status` ENUM ('hadir', 'tidak_hadir'),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp NULL,
  `deleted_at` timestamp NULL
);

CREATE TABLE `photos` (
  `id` varchar(36) PRIMARY KEY NOT NULL,
  `documentation_id` varchar(36) NOT NULL,
  `url` varchar(36),
  `name` varchar(36),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp NULL,
  `deleted_at` timestamp NULL
); 

ALTER TABLE `events` ADD FOREIGN KEY (`branch_id`) REFERENCES `branchs` (`id`);

ALTER TABLE `documentations` ADD FOREIGN KEY (`branch_id`) REFERENCES `branchs` (`id`);

ALTER TABLE `participants` ADD FOREIGN KEY (`event_id`) REFERENCES `events` (`id`);

ALTER TABLE `photos` ADD FOREIGN KEY (`documentation_id`) REFERENCES `documentations` (`id`);


-- migrate:down
DROP TABLE `photos`;
DROP TABLE `participants`;
DROP TABLE `documentations`;
DROP TABLE `events`;
DROP TABLE `branchs`;
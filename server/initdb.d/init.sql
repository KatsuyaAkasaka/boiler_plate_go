CREATE DATABASE IF NOT EXISTS boiler_plate_go_local;

DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
  `user_id` varchar(255),
  `uuid` varchar(255) NOT NULL,
  `nick_name` varchar(255) NOT NULL,
  `profile_image_uri` varchar(255),
  `email` varchar(255) NOT NULL,
  `description` varchar(255),
  `social_link` varchar(255),
  `gender` tinyint UNSIGNED,
  `is_official` boolean default false,
  `send_mail_status` tinyint,
  `customer_id` varchar(255) UNIQUE,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  `deleted_at` timestamp,
  PRIMARY KEY (user_id),
  UNIQUE KEY (uuid),
  UNIQUE KEY (email),
  UNIQUE KEY (customer_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 取得失敗時に0埋めされるので初期値は必ず1にする
SET SESSION auto_increment_offset = 1;
SET SESSION auto_increment_increment = 1;

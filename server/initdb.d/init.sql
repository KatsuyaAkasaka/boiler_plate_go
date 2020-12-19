DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
  `user_id` varchar(255),
  `nick_name` varchar(255) NOT NULL,
  `profile_image_uri` varchar(255),
  `email` varchar(255) UNIQUE NOT NULL,
  `description` varchar(255),
  `social_link` varchar(255),
  `gender` tinyint UNSIGNED,
  `identify_status` tinyint UNSIGNED,
  `customer_id` varchar(255) UNIQUE,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  `deleted_at` timestamp,
  PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

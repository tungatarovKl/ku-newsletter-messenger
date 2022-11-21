use messenger;
CREATE TABLE `users` (
 `id` int unsigned primary key NOT NULL AUTO_INCREMENT,
  `Telegram_Id` bigint UNSIGNED NOT NULL,
  `First_Name` varchar(150) DEFAULT NULL,
  `Last_Name` varchar(150) DEFAULT NULL,
  `Chat_Id` bigint UNSIGNED UNIQUE NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
  );


CREATE TABLE `tokens` (
  `id` int unsigned primary key NOT NULL AUTO_INCREMENT,
  `Token_Key` varchar(150) UNIQUE NOT NULL,
  `Service_Name` varchar(150) NOT NULL,
    `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
  );
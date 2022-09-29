CREATE TABLE IF NOT EXISTS`User` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `fullName` VARCHAR(255) NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  `gender` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(20) NOT NULL,
  `blocked` BOOL NOT NULL,
  `datecreated` DATETIME NOT NULL,
  `dateupdated` DATETIME
);

CREATE TABLE IF NOT EXISTS `User_Password` (
  `id` INT PRIMARY KEY NOT NULL,
  `user_id` INT UNIQUE,
  `password` VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `Tutor` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `fullName` VARCHAR(255) NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  `gender` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(20) NOT NULL,
  `validate` VARCHAR(255) NOT NULL,
  `adminId` INT,
  `datecreated` DATETIME NOT NULL,
  `dateupdated` DATETIME
);

CREATE TABLE IF NOT EXISTS `Tutor_Password` (
  `id` INT PRIMARY KEY NOT NULL,
  `tutor_id` INT UNIQUE,
  `password` VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `Admin` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `datecreated` DATETIME NOT NULL,
  `dateupdate` DATETIME NOT NULL
);

CREATE UNIQUE INDEX `User_index_0` ON `User` (`username`);

CREATE UNIQUE INDEX `User_index_1` ON `User` (`email`);

CREATE UNIQUE INDEX `User_index_2` ON `User` (`phone`);

CREATE UNIQUE INDEX `Tutor_index_3` ON `Tutor` (`username`);

CREATE UNIQUE INDEX `Tutor_index_4` ON `Tutor` (`email`);

CREATE UNIQUE INDEX `Tutor_index_5` ON `Tutor` (`phone`);

ALTER TABLE `User_Password` ADD FOREIGN KEY (`user_id`) REFERENCES `User` (`id`);

ALTER TABLE `Tutor` ADD FOREIGN KEY (`adminId`) REFERENCES `Admin` (`id`);

ALTER TABLE `Tutor_Password` ADD FOREIGN KEY (`tutor_id`) REFERENCES `Tutor` (`id`);

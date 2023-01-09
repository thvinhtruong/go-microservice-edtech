CREATE TABLE IF NOT EXISTS`User` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `fullName` VARCHAR(255) NOT NULL,
  `gender` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(15) NOT NULL,
  `datecreated` DATETIME NOT NULL,
  `dateupdated` DATETIME
);

CREATE TABLE IF NOT EXISTS `User_Password` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` INT UNIQUE NOT NULL,
  `password` VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `Tutor` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `fullName` VARCHAR(100) NOT NULL,
  `gender` VARCHAR(6) NOT NULL,
  `age` INT(2) NOT NULL,
  `phone` VARCHAR(10) NOT NULL,
  `topic` VARCHAR(10) NOT NULL,
  `country` VARCHAR(10) NOT NULL,
  `city` VARCHAR(15) NOT NULL,
  `datecreated` DATETIME NOT NULL,
  `dateupdated` DATETIME,
  KEY `Tutor_search_index` (`gender`,`topic`,`country`,`city`,`age`)
);

CREATE TABLE IF NOT EXISTS `Tutor_Password` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `tutor_id` INT UNIQUE NOT NULL,
  `password` VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `Admin` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `datecreated` DATETIME NOT NULL,
  `dateupdate` DATETIME NOT NULL
);

CREATE UNIQUE INDEX `User_index_1` ON `User` (`phone`);

CREATE UNIQUE INDEX `Tutor_index_1` ON `Tutor` (`phone`);

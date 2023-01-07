CREATE TABLE IF NOT EXISTS`Course` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `price` INT NOT NULL,
  `datecreated` DATETIME NOT NULL,
  `dateupdated` DATETIME
);

CREATE TABLE IF NOT EXISTS `Course_Tutor` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `course_id` INT NOT NULL,
  `tutor_id` INT NOT NULL,
  INDEX `course_id_tutor_id` (`course_id`, `tutor_id`)
);

CREATE TABLE IF NOT EXISTS `Course_User` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `course_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  INDEX `course_id_user_id` (`course_id`, `user_id`)
);

CREATE TABLE IF NOT EXISTS `Feedback` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `course_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `content` VARCHAR(255) NOT NULL,
  `datecreated` DATETIME NOT NULL,
  `dateupdated` DATETIME,
  INDEX `course_id_user_id` (`course_id`, `user_id`)
);

CREATE TABLE IF NOT EXISTS `Lecture` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `course_id` INT NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `content` VARCHAR(255) NOT NULL,
  `datecreated` DATETIME NOT NULL,
  `dateupdated` DATETIME
);


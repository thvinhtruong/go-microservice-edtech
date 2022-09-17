CREATE DATABASE IF NOT EXISTS SepFirst;
use SepFirst;

CREATE TABLE Student (
	id INT NOT NULL AUTO_INCREMENT,
	fullName VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
	gender VARCHAR(255) NOT NULL,
	email VARCHAR(255),
	blocked BOOL NOT NULL,
	datecreated DATETIME NOT NULL,
	dateupdated DATETIME,
	UNIQUE (username),
	UNIQUE(email),
	PRIMARY KEY (id)
);

CREATE TABLE Student_Password (
    id INT NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE Tutor (
	id INT NOT NULL AUTO_INCREMENT,
    fullName VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
	gender VARCHAR(255) NOT NULL,
	email VARCHAR(255),
	validate VARCHAR(255) NOT NULL,
	datecreated DATETIME NOT NULL,
	dateupdated DATETIME,
	UNIQUE (username),
	PRIMARY KEY (id)
);

CREATE TABLE Tutor_Password(
    id INT NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE Admin (
	id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(username),
);

CREATE TABLE Admin_Password (
    id INT NOT NULL,
    Password VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);
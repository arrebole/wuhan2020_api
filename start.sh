#! /usr/bash

#CREATE TABLE `archives`(
#   `id` INT UNSIGNED AUTO_INCREMENT,
#   `province` VARCHAR(255),
#   `city` VARCHAR(255),
#   `publish_time` VARCHAR(255),
#   `publish_date` VARCHAR(255),
#   `title` VARCHAR(255),
#   `content` TEXT,
#   `link` TEXT,
#   `links_to_pic` TEXT,
#   `announce_type` VARCHAR(100),
#   PRIMARY KEY ( `id` )
#)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#CREATE TABLE `logs`(
#   `id` INT UNSIGNED AUTO_INCREMENT,
#   `time` VARCHAR(255),
#   `uploader` VARCHAR(255),
#   `province` VARCHAR(255),
#   `city` VARCHAR(255),
#   PRIMARY KEY ( `id` )
#)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
PATH=/home/godalike/go1.13.8/bin:$PATH
go build -o main.out ./main.go
read -p "Enter database password: " password
export PASSWORD=$password
(./main.out 1>./var/access.log 2>&1 &)



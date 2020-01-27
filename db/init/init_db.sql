CREATE DATABASE test_db character SET utf8mb4 collate utf8mb4_bin;

USE test_db;

CREATE TABLE users (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(255) COLLATE utf8mb4_bin NOT NULL,
  age int(11) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

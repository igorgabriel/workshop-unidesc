CREATE DATABASE workshop CHARSET = Latin1 COLLATE = latin1_swedish_ci;

USE workshop;

CREATE TABLE workshop (
  id int PRIMARY KEY AUTO_INCREMENT,
  nome text,
  palestrante int
);
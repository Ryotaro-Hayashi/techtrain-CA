CREATE DATABASE IF NOT EXISTS dojo_db;

USE dojo_db;

CREATE TABLE users(
  id              INTEGER PRIMARY KEY AUTO_INCREMENT,
  name            VARCHAR(40) NOT NULL,
  token           VARCHAR(255),
  created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO users (name, token) VALUES ("林 遼太朗", "11111");
INSERT INTO users (name, token) VALUES ("John Titor", "22222");

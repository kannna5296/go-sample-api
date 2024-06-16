CREATE DATABASE sample;

USE sample;
CREATE TABLE book (
  id VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  release_date DATETIME,
  PRIMARY KEY CLUSTERED (
    id ASC
  )
);

CREATE TABLE rental_user (
  id VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  phone VARCHAR(255),
  mail VARCHAR(255) NOT NULL,
  PRIMARY KEY CLUSTERED (
    id ASC
  )
);

CREATE TABLE rental (
  id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  book_id VARCHAR(255) NOT NULL,
  deadline DATETIME NOT NULL,
  is_returned BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY CLUSTERED (
    id ASC
  )
);

INSERT INTO book VALUES('1', 'The Fellowship of the Ring', 'J. R. R. Tolkien', '1954-07-29');
INSERT INTO book VALUES('2', 'The Two Towers', 'J. R. R. Tolkien', '1954-11-11');
INSERT INTO book VALUES('3', 'The Return of the King', 'J. R. R. Tolkien', '1955-10-20');

INSERT INTO rental_user (id, name, phone, mail)
  VALUES
('1','tom', '000-0000', 'example1@gmail.com'),
('2','jerry', '000-0001', 'example2@gmail.com');

INSERT INTO rental (id, book_id, user_id, deadline, is_returned)
 VALUES
('1', '1', '1', '2024-01-08', 1),
('2', '2', '1', '2024-01-09', 0),
('3', '3', '1', '2024-01-10', 0);-- 未返却
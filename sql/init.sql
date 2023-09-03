CREATE DATABASE sample;

USE sample;
CREATE TABLE sample.book (
  id VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  release_date DATETIME,
  PRIMARY KEY CLUSTERED (
    id ASC
  )
);

CREATE TABLE sample.rental (
  id VARCHAR(255) NOT NULL,
  user_id INT NOT NULL,
  book_id VARCHAR(255) NOT NULL,
  rental_date DATETIME NOT NULL,
  return_deadline DATETIME NOT NULL,
  PRIMARY KEY CLUSTERED (
    id ASC
  )
);

INSERT INTO book VALUES('1', 'The Fellowship of the Ring', 'J. R. R. Tolkien', '1954-07-29');
INSERT INTO book VALUES('2', 'The Two Towers', 'J. R. R. Tolkien', '1954-11-11');
INSERT INTO book VALUES('3', 'The Return of the King', 'J. R. R. Tolkien', '1955-10-20');
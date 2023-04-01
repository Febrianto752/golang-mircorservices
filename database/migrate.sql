-- DDL
CREATE DATABASE golang_dasar;

CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(50) NOT NULL,
  author VARCHAR(50) NOT NULL,
  description TEXT NOT NULL
);


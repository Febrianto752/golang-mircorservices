-- DDL
CREATE DATABASE golang_dasar;

CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  name_book VARCHAR(50) NOT NULL,
  author VARCHAR(50) NOT NULL,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
);


-- +goose Up
-- +goose StatementBegin
CREATE TABLE authors (
  id text NOT NULL PRIMARY KEY,
  name text,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE books (
  id text NOT NULL PRIMARY KEY,
  title text ,
  author_id text REFERENCES authors(id),
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
DROP TABLE authors;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE authors (
  id text NOT NULL PRIMARY KEY,
  name text UNIQUE NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE books (
  id text NOT NULL PRIMARY KEY,
  title text ,
  author_id text NOT NULL REFERENCES authors(id),
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO authors (id, name) VALUES ('7c1f7bb6-113b-47b0-997a-d4a1994ba8b0', 'Kevin Yobeth');
INSERT INTO authors (id, name) VALUES ('51642ead-b935-4421-b39b-eca86707fca7', 'Kimi');

INSERT INTO books (id, title, author_id) VALUES ('acea1ac3-467b-45f5-8acd-1b57f7ad8fda', 'Bobah: The Story', '7c1f7bb6-113b-47b0-997a-d4a1994ba8b0');
INSERT INTO books (id, title, author_id) VALUES ('4c87ca84-f000-4cd5-b92e-b65f09b0b479', 'Ayam Penyet Bu Ria', '7c1f7bb6-113b-47b0-997a-d4a1994ba8b0');
INSERT INTO books (id, title, author_id) VALUES ('344b145e-39c8-476f-bad7-c636bdda5ed0', 'The Visual MBA', '51642ead-b935-4421-b39b-eca86707fca7');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
DROP TABLE authors;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS books (
  id  UUID NOT NULL PRIMARY KEY,
  title VARCHAR(50) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  created_by VARCHAR(50),
  updated_at TIMESTAMP,
  updated_by VARCHAR(50),
  deleted_at TIMESTAMP,
  deleted_by VARCHAR(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS books;
-- +goose StatementEnd


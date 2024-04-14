-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS author_book (
  author_id UUID NOT NULL REFERENCES authors(id),
  book_id UUID NOT NULL REFERENCES books(id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  created_by VARCHAR(50),
  updated_at TIMESTAMP,
  updated_by VARCHAR(50),
  deleted_at TIMESTAMP,
  deleted_by VARCHAR(50),

  PRIMARY KEY (author_id, book_id)
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS author_book;
-- +goose StatementEnd

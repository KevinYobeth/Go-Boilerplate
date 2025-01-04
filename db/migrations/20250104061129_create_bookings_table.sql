-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS bookings (
  id UUID NOT NULL PRIMARY KEY,
  book_id UUID NOT NULL REFERENCES books(id),
  date_from DATE NOT NULL,
  date_to DATE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  created_by VARCHAR(50),
  updated_at TIMESTAMP,
  updated_by VARCHAR(50),
  deleted_at TIMESTAMP,
  deleted_by VARCHAR(50)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS bookings;
-- +goose StatementEnd

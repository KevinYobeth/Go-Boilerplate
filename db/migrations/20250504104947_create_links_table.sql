-- +goose Up
-- +goose StatementBegin
CREATE TABLE links (
  id UUID NOT NULL PRIMARY KEY,
  slug VARCHAR(100) NOT NULL,
  url TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP,
  updated_by VARCHAR(50),
  deleted_at TIMESTAMP,
  deleted_by VARCHAR(50)
);

CREATE TABLE link_visit_events (
  id UUID NOT NULL PRIMARY KEY,
  link_id UUID NOT NULL REFERENCES links(id),
  ip_address INET,
  user_agent TEXT,
  referer TEXT,
  country_code VARCHAR(2),
  device_type VARCHAR(50),
  browser VARCHAR(50),
  visited_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);  

CREATE INDEX idx_links_slug ON links(slug);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS link_visit_events;
DROP TABLE IF EXISTS links;

DROP INDEX IF EXISTS idx_links_slug;
-- +goose StatementEnd

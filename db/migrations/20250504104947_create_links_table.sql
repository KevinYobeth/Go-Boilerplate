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

CREATE TABLE link_visits (
  id UUID NOT NULL PRIMARY KEY,
  link_id UUID NOT NULL REFERENCES links(id),
  slug VARCHAR(100) NOT NULL,
  ip_address INET,
  user_agent TEXT,
  visited_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE link_visit_snapshots (
  id UUID NOT NULL PRIMARY KEY,
  link_id UUID NOT NULL REFERENCES links(id),
  total INT NOT NULL,
  last_snapshot_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_links_slug ON links(slug);
CREATE INDEX idx_link_visits_link_id ON link_visits(link_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS link_visit_snapshots;
DROP TABLE IF EXISTS link_visits;
DROP TABLE IF EXISTS links;

DROP INDEX IF EXISTS idx_links_slug;
DROP INDEX IF EXISTS idx_link_visits_link_id;
-- +goose StatementEnd

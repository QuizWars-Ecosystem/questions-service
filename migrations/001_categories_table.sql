-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS categories (
    id SMALLSERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_categories_name ON categories(name);

---- create above / drop below ----

DROP INDEX IF EXISTS idx_unique_categories_name;

DROP TABLE IF EXISTS categories;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

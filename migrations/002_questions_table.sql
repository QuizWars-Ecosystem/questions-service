-- Write your migrate up statements here

CREATE TYPE difficulty_type AS ENUM ('easy', 'medium', 'hard', 'very hard');
CREATE TYPE type_enum AS ENUM ('single', 'multi', 'betting');
CREATE TYPE source_enum AS ENUM ('text', 'image', 'audio', 'animation', 'video');
CREATE TYPE language_enum AS ENUM ('eng', 'rus');

CREATE TABLE IF NOT EXISTS questions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    text VARCHAR(500) NOT NULL,
    text_hash CHAR(32) UNIQUE NOT NULL,
    category_id INTEGER NOT NULL REFERENCES categories(id),
    type type_enum NOT NULL DEFAULT 'single' CHECK ( type IN ('single', 'multi', 'betting')),
    source source_enum NOT NULL DEFAULT 'text' CHECK ( source IN ('text', 'image', 'audio', 'animation', 'video')),
    difficulty difficulty_type NOT NULL DEFAULT 'easy' CHECK (difficulty IN ('easy', 'medium', 'hard', 'very hard')),
    language language_enum NOT NULL DEFAULT 'eng' CHECK (language IN ('eng', 'rus')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_questions_type ON questions(type);
CREATE INDEX IF NOT EXISTS idx_questions_source ON questions(source);
CREATE INDEX IF NOT EXISTS idx_questions_difficulty ON questions(difficulty);
CREATE INDEX IF NOT EXISTS idx_questions_language ON questions(language);
CREATE INDEX IF NOT EXISTS idx_questions_created_at ON questions(created_at);
CREATE INDEX IF NOT EXISTS idx_questions_category_id ON questions(category_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_questions_hash ON questions(text_hash);

---- create above / drop below ----

DROP INDEX IF EXISTS idx_unique_questions_hash;
DROP INDEX IF EXISTS idx_questions_category_id;
DROP INDEX IF EXISTS idx_questions_created_at;
DROP INDEX IF EXISTS idx_questions_language;
DROP INDEX IF EXISTS idx_questions_difficulty;
DROP INDEX IF EXISTS idx_questions_source;
DROP INDEX IF EXISTS idx_questions_type;

DROP TABLE IF EXISTS questions;

DROP TYPE IF EXISTS language_enum;
DROP TYPE IF EXISTS type_enum;
DROP TYPE IF EXISTS difficulty_type;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

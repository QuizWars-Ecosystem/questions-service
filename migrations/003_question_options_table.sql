-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS question_options (
    id UUID DEFAULT gen_random_uuid(),
    question_id UUID NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    text VARCHAR(256) NOT NULL,
    is_correct BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (id, question_id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS question_options;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

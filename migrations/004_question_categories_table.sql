-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS question_categories (
    question_id UUID NOT NULL REFERENCES questions(id),
    category_id SMALLINT NOT NULL REFERENCES categories(id),
    PRIMARY KEY (question_id, category_id)
);

CREATE INDEX IF NOT EXISTS idx_question_categories_question_id ON question_categories(question_id);
CREATE INDEX IF NOT EXISTS idx_question_categories_category_id ON question_categories(category_id);

---- create above / drop below ----

DROP INDEX IF EXISTS idx_question_categories_category_id;
DROP INDEX IF EXISTS idx_question_categories_category_id;

DROP TABLE IF EXISTS question_categories;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

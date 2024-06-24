-- Write your migrate up statements here
CREATE OR REPLACE FUNCTION validate_vocabulary_example() RETURNS TRIGGER AS $$
BEGIN
    -- Validate translated JSONB
    IF NOT (
        jsonb_typeof(NEW.translated) = 'object'
    ) THEN
        RAISE EXCEPTION 'Invalid translated JSONB structure';
END IF;

    -- Validate main_word JSONB
    IF NOT (
        jsonb_typeof(NEW.main_word->'word') = 'string' AND
        jsonb_typeof(NEW.main_word->'base') = 'string' AND
        jsonb_typeof(NEW.main_word->'pos') = 'string' AND
        jsonb_typeof(NEW.main_word->'definition') = 'string' AND
        jsonb_typeof(NEW.main_word->'translated') = 'object'
    ) THEN
        RAISE EXCEPTION 'Invalid main_word JSONB structure';
END IF;

    -- Validate pos_tags JSONB
    IF NOT (
        jsonb_typeof(NEW.pos_tags) = 'array'
    ) THEN
        RAISE EXCEPTION 'Invalid pos_tags JSONB structure';
END IF;

    -- Validate sentiment JSONB
    IF NOT (
        jsonb_typeof(NEW.sentiment->'polarity') = 'number' AND
        jsonb_typeof(NEW.sentiment->'subjectivity') = 'number'
    ) THEN
        RAISE EXCEPTION 'Invalid sentiment JSONB structure';
END IF;

    -- Validate dependencies JSONB
    IF NOT (
        jsonb_typeof(NEW.dependencies) = 'array'
    ) THEN
        RAISE EXCEPTION 'Invalid dependencies JSONB structure';
END IF;

    -- Validate verbs JSONB
    IF NOT (
        jsonb_typeof(NEW.verbs) = 'array'
    ) THEN
        RAISE EXCEPTION 'Invalid verbs JSONB structure';
END IF;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER validate_vocabulary_example_trigger
    BEFORE INSERT OR UPDATE ON vocabulary_examples
    FOR EACH ROW EXECUTE FUNCTION validate_vocabulary_example();


---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

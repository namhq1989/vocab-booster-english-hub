-- Write your migrate up statements here

-- VOCABULARY --

CREATE OR REPLACE FUNCTION validate_vocabulary() RETURNS TRIGGER AS $$
BEGIN
    -- Validate definitions JSONB
    IF NOT (
        jsonb_typeof(NEW.definitions) = 'array'
    ) THEN
        RAISE EXCEPTION 'Invalid definitions JSONB structure';
END IF;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER validate_vocabulary_trigger
    BEFORE INSERT OR UPDATE ON vocabularies
    FOR EACH ROW EXECUTE FUNCTION validate_vocabulary();

-- VOCABULARY EXAMPLE --

CREATE OR REPLACE FUNCTION validate_vocabulary_example() RETURNS TRIGGER AS $$
BEGIN
    -- Validate content JSONB
    IF NOT (
        jsonb_typeof(NEW.content) = 'object'
    ) THEN
        RAISE EXCEPTION 'Invalid content JSONB structure';
END IF;

    -- Validate main_word JSONB
    IF NOT (
        jsonb_typeof(NEW.main_word->'word') = 'string' AND
        jsonb_typeof(NEW.main_word->'base') = 'string' AND
        jsonb_typeof(NEW.main_word->'pos') = 'string' AND
        jsonb_typeof(NEW.main_word->'definition') = 'object'
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

-- COMMUNITY SENTENCE --

CREATE OR REPLACE FUNCTION validate_community_sentence() RETURNS TRIGGER AS $$
BEGIN
    -- Validate content JSONB
    IF NOT (
        jsonb_typeof(NEW.content) = 'object'
    ) THEN
        RAISE EXCEPTION 'Invalid content JSONB structure';
END IF;

    -- Validate sentiment JSONB
    IF NOT (
        jsonb_typeof(NEW.sentiment->'polarity') = 'number' AND
        jsonb_typeof(NEW.sentiment->'subjectivity') = 'number'
    ) THEN
        RAISE EXCEPTION 'Invalid sentiment JSONB structure';
END IF;

    -- Validate clauses JSONB
    IF NOT (
        jsonb_typeof(NEW.clauses) = 'array'
    ) THEN
        RAISE EXCEPTION 'Invalid clauses JSONB structure';
END IF;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER validate_community_sentence_trigger
    BEFORE INSERT OR UPDATE ON community_sentences
    FOR EACH ROW EXECUTE FUNCTION validate_community_sentence();

-- COMMUNITY SENTENCE DRAFT --

CREATE OR REPLACE FUNCTION validate_community_sentence_draft() RETURNS TRIGGER AS $$
BEGIN
    -- Validate content JSONB
    IF NOT (
        jsonb_typeof(NEW.content) = 'object'
    ) THEN
        RAISE EXCEPTION 'Invalid content JSONB structure';
END IF;

    -- Validate sentiment JSONB
    IF NOT (
        jsonb_typeof(NEW.sentiment->'polarity') = 'number' AND
        jsonb_typeof(NEW.sentiment->'subjectivity') = 'number'
    ) THEN
        RAISE EXCEPTION 'Invalid sentiment JSONB structure';
END IF;

    -- Validate clauses JSONB
    IF NOT (
        jsonb_typeof(NEW.clauses) = 'array'
    ) THEN
        RAISE EXCEPTION 'Invalid clauses JSONB structure';
END IF;

    -- Validate grammar errors JSONB
    IF NOT (
        jsonb_typeof(NEW.grammar_errors) = 'array'
    ) THEN
        RAISE EXCEPTION 'Invalid grammar errors JSONB structure';
END IF;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER validate_community_sentence_draft_trigger
    BEFORE INSERT OR UPDATE ON community_sentence_drafts
    FOR EACH ROW EXECUTE FUNCTION validate_community_sentence_draft();

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

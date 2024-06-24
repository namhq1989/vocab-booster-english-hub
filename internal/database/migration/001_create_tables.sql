-- Write your migrate up statements here

CREATE SCHEMA IF NOT EXISTS public;

-- VOCABULARY --

CREATE TABLE vocabularies (
    id TEXT PRIMARY KEY,
    author_id TEXT NOT NULL,
    term VARCHAR(30) NOT NULL,
    parts_of_speech TEXT[] NOT NULL,
    ipa VARCHAR(30) NOT NULL,
    audio VARCHAR(50) NOT NULL,
    synonyms TEXT[] NOT NULL DEFAULT '{}',
    antonyms TEXT[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE TABLE vocabulary_examples (
    id TEXT PRIMARY KEY,
    vocabulary_id TEXT NOT NULL,
    audio VARCHAR(50) NOT NULL,
    content TEXT NOT NULL,
    translated JSONB NOT NULL,
    main_word JSONB NOT NULL,
    pos_tags JSONB NOT NULL DEFAULT '{}',
    sentiment JSONB NOT NULL,
    dependencies JSONB NOT NULL DEFAULT '{}',
    verbs JSONB NOT NULL DEFAULT '{}',
    level VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    CONSTRAINT fk_vocabulary
        FOREIGN KEY(vocabulary_id)
        REFERENCES vocabularies(id)
);

CREATE TABLE verb_conjugations (
    id TEXT PRIMARY KEY,
    vocabulary_id TEXT NOT NULL,
    value VARCHAR(30) NOT NULL,
    base VARCHAR(30) NOT NULL,
    form VARCHAR(30) NOT NULL,

    CONSTRAINT fk_vocabulary
        FOREIGN KEY(vocabulary_id)
        REFERENCES vocabularies(id),

    CONSTRAINT unique_value_form
        UNIQUE (value, form)
);

CREATE TABLE vocabulary_scraping_items (
    id TEXT PRIMARY KEY,
    term VARCHAR(30) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    CONSTRAINT unique_term
        UNIQUE (term)
);

-- EXERCISE --

CREATE TABLE exercises (
    id TEXT PRIMARY KEY,
    vocabulary_example_id TEXT NOT NULL,
    audio VARCHAR(50) NOT NULL,
    level VARCHAR(20) NOT NULL,
    content TEXT NOT NULL,
    translated JSON NOT NULL,
    vocabulary VARCHAR(30) NOT NULL,
    correct_answer VARCHAR(30) NOT NULL,
    options TEXT[] NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    CONSTRAINT fk_vocabulary_example
        FOREIGN KEY(vocabulary_example_id)
        REFERENCES vocabulary_examples(id)
);

CREATE TABLE user_exercise_statuses (
    id TEXT PRIMARY KEY,
    exercise_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    correct_streak INT NOT NULL,
    is_favorite BOOLEAN NOT NULL,
    is_mastered BOOLEAN NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    next_review_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_exercise
        FOREIGN KEY(exercise_id)
        REFERENCES exercises(id)
);

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
DROP TABLE exercises CASCADE;
DROP TABLE user_exercise_statuses CASCADE;
